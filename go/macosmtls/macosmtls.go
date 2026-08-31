package main

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"time"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/stream"
	"google.golang.org/protobuf/types/known/timestamppb"
	"namespacelabs.dev/integrations/api"
	"namespacelabs.dev/integrations/api/builds"
	"namespacelabs.dev/integrations/api/compute"
	"namespacelabs.dev/integrations/auth"
	"namespacelabs.dev/integrations/auth/nstls"
	"namespacelabs.dev/integrations/examples"
	computepb "namespacelabs.dev/integrations/proto/namespace/cloud/compute/v1beta"
)

const (
	ingressName = "echo"
	servicePort = 15000
)

var basedir = flag.String("basedir", "", "If not specified, it's computed from the binary's location.")

func main() {
	flag.Parse()

	if err := do(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func do(ctx context.Context) error {
	basedir, err := examples.ComputeBaseDir(*basedir)
	if err != nil {
		return err
	}

	token, err := auth.LoadDefaults()
	if err != nil {
		return err
	}

	dir, err := os.MkdirTemp("", "macosmtls")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	target := filepath.Join(dir, "entrypoint")
	if err := gobuild(ctx, target, filepath.Join(basedir, "echoserver")); err != nil {
		return err
	}

	imageRef, err := pushImage(ctx, token, dir)
	if err != nil {
		return err
	}

	endpoint, err := runInstance(ctx, os.Stderr, token, imageRef)
	if err != nil {
		return err
	}

	callCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return callService(callCtx, auth.TenantCertificateSource(token), endpoint)
}

func gobuild(ctx context.Context, target, srcdir string) error {
	fmt.Fprintf(os.Stderr, "Running: go build -v -o %s .\n", target)

	cmd := exec.CommandContext(ctx, "go", "build", "-v", "-o", target, ".")
	cmd.Dir = srcdir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(slices.Clone(os.Environ()), "CGO_ENABLED=0", "GOOS=darwin", "GOARCH=arm64")
	return cmd.Run()
}

func pushImage(ctx context.Context, token api.TokenSource, dir string) (string, error) {
	var tarBytes bytes.Buffer
	w := tar.NewWriter(&tarBytes)
	if err := w.AddFS(os.DirFS(dir)); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	image, err := mutate.AppendLayers(empty.Image, stream.NewLayer(io.NopCloser(bytes.NewReader(tarBytes.Bytes()))))
	if err != nil {
		return "", fmt.Errorf("produce image: %w", err)
	}

	repository, err := builds.NSCRImage(ctx, token, "example/macosmtls/echoserver")
	if err != nil {
		return "", fmt.Errorf("compute repository: %w", err)
	}
	parsed, err := name.NewTag(repository)
	if err != nil {
		return "", fmt.Errorf("parse image reference: %w", err)
	}
	if err := remote.Write(parsed, image, remote.WithContext(ctx), remote.WithAuthFromKeychain(builds.NewNSCRKeychain(token))); err != nil {
		return "", fmt.Errorf("push image: %w", err)
	}

	digest, err := image.Digest()
	if err != nil {
		return "", fmt.Errorf("compute digest: %w", err)
	}
	return parsed.Digest(digest.String()).String(), nil
}

func runInstance(ctx context.Context, debugLog io.Writer, token api.TokenSource, imageRef string) (string, error) {
	client, err := compute.NewClient(ctx, token)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Compute.CreateInstance(ctx, &computepb.CreateInstanceRequest{
		Shape: &computepb.InstanceShape{
			VirtualCpu:      6,
			MemoryMegabytes: 14 * 1024,
			Os:              "macos",
			MachineArch:     "arm64",
		},
		DocumentedPurpose: "macosmtls example",
		Deadline:          timestamppb.New(time.Now().Add(30 * time.Minute)),
		Applications: []*computepb.ApplicationRequest{{
			Name:     "echoserver",
			ImageRef: imageRef,
			Command:  "entrypoint",
		}},
		Ingresses: []*computepb.Ingress{{
			Name:                 ingressName,
			Mode:                 computepb.Ingress_TCP,
			Port:                 servicePort,
			ClientAuthentication: computepb.IngressClientAuthentication_INGRESS_CLIENT_AUTHENTICATION_MTLS,
		}},
	})
	if err != nil {
		return "", err
	}

	var endpoint string
	for _, ingress := range resp.GetExtendedMetadata().GetTlsBackedPort() {
		if ingress.GetName() == ingressName {
			endpoint = ingress.GetServerName()
			break
		}
	}
	if endpoint == "" {
		return "", fmt.Errorf("create response did not include the %q ingress endpoint", ingressName)
	}

	fmt.Fprintf(debugLog, "[namespace] Instance created: %s\n", resp.InstanceUrl)
	fmt.Fprintf(debugLog, "[namespace] mTLS endpoint: %s\n", endpoint)
	fmt.Fprintln(debugLog, "[namespace] Waiting until instance becomes ready")
	if _, err := client.Compute.WaitInstanceSync(ctx, &computepb.WaitInstanceRequest{
		InstanceId: resp.Metadata.InstanceId,
	}); err != nil {
		return "", err
	}

	return endpoint, nil
}

func callService(ctx context.Context, certificates api.CertificateSource, endpoint string) error {
	dialer := tls.Dialer{Config: nstls.ClientConfig(ctx, certificates)}
	conn, err := dialer.DialContext(ctx, "tcp", endpoint)
	if err != nil {
		return fmt.Errorf("dial %s: %w", endpoint, err)
	}
	defer conn.Close()

	if _, err := io.WriteString(conn, "hello from the mTLS client\n"); err != nil {
		return err
	}
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Printf("received: %q\n", reply)
	return nil
}
