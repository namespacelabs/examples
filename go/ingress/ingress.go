package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	computepb "buf.build/gen/go/namespace/cloud/protocolbuffers/go/proto/namespace/cloud/compute/v1beta"
	"google.golang.org/protobuf/types/known/timestamppb"
	"namespacelabs.dev/integrations/api"
	"namespacelabs.dev/integrations/api/compute"
	"namespacelabs.dev/integrations/auth"
)

var deadline = flag.Duration("deadline", 5*time.Minute, "How long the instance should run for.")

func main() {
	flag.Parse()

	token, err := auth.LoadDefaults()
	if err != nil {
		log.Fatal(err)
	}

	if err := run(context.Background(), token); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, token api.TokenSource) error {
	cli, err := compute.NewClient(ctx, token)
	if err != nil {
		return err
	}

	defer cli.Close()

	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")

	// Create an instance with an nginx container on the host network, listening on port 8080.
	resp, err := cli.Compute.CreateInstance(ctx, &computepb.CreateInstanceRequest{
		Shape: &computepb.InstanceShape{
			VirtualCpu:      2,
			MemoryMegabytes: 4 * 1024,
			MachineArch:     "amd64",
		},
		DocumentedPurpose: "ingress example",
		Deadline:          timestamppb.New(time.Now().Add(*deadline)),
		Containers: []*computepb.ContainerRequest{{
			Name:     "nginx",
			ImageRef: "nginx",
			Args:     []string{"sh", "-c", "echo 'server { listen 8080; location / { default_type text/plain; return 200 \"hello from nginx\\n\"; } }' > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'"},
			Network: computepb.ContainerRequest_HOST,
		}},
	})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "[namespace] Instance: %s\n", resp.InstanceUrl)

	_ = enc.Encode(resp)

	// Wait until the instance is ready.
	fmt.Fprintf(os.Stderr, "[Waiting until instance becomes ready]\n")

	_, err = cli.Compute.WaitInstanceSync(ctx, &computepb.WaitInstanceRequest{
		InstanceId: resp.Metadata.InstanceId,
	})
	if err != nil {
		return err
	}

	// Expose port 8080 via public ingress.
	ingressResp, err := cli.Compute.CreateIngress(ctx, &computepb.CreateIngressRequest{
		InstanceId: resp.Metadata.InstanceId,
		Ingresses: []*computepb.IngressRequest{{
			Name: "nginx",
			ExportedPortBackend: &computepb.ExportedPortBackend{
				Port: 8080,
			},
		}},
	})
	if err != nil {
		return err
	}

	for _, ingress := range ingressResp.AllocatedIngresses {
		fmt.Fprintf(os.Stdout, "https://%s\n", ingress.Fqdn)
	}

	return nil
}
