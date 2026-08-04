package main

import (
	"context"
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

var (
	uniqueTag = flag.String("unique_tag", "ensure-with-cache", "Unique tag used to create or reuse the instance.")
	cacheTag  = flag.String("cache_tag", "ensure-with-cache-data", "Tag used to restore the cache volume across instance lifetimes.")
	deadline  = flag.Duration("deadline", time.Hour, "How long a newly created instance should run.")
)

func main() {
	flag.Parse()

	token, err := auth.LoadDefaults()
	if err != nil {
		log.Fatal(err)
	}

	endpoint, err := ensure(context.Background(), token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "https://%s\n", endpoint)
}

func ensure(ctx context.Context, token api.TokenSource) (string, error) {
	cli, err := compute.NewClient(ctx, token)
	if err != nil {
		return "", err
	}
	defer cli.Close()

	// Repeating this request with the same unique tag returns the running instance instead of creating another one.
	resp, err := cli.Compute.CreateInstance(ctx, &computepb.CreateInstanceRequest{
		Shape: &computepb.InstanceShape{
			VirtualCpu:      2,
			MemoryMegabytes: 4 * 1024,
			MachineArch:     "amd64",
			Os:              "linux",
		},
		DocumentedPurpose: "ensurewithcache example",
		Deadline:          timestamppb.New(time.Now().Add(*deadline)),
		Experimental: &computepb.CreateInstanceRequest_ExperimentalFeatures{
			UniqueTag: *uniqueTag,
		},
		Containers: []*computepb.ContainerRequest{{
			Name:       "server",
			ImageRef:   "alpine:3.21",
			Entrypoint: []string{"/bin/sh", "-c"},
			Args: []string{`if [ ! -f /cache/index.html ]; then
  printf 'cache initialized by %s at %s\n' "$HOSTNAME" "$(date -u +%FT%TZ)" > /cache/index.html
fi
exec httpd -f -p 8080 -h /cache`},
			// A later instance can best-effort restore the latest committed volume carrying this cache tag.
			Volumes: []*computepb.VolumeRequest{{
				MountPoint:      "/cache",
				Tag:             *cacheTag,
				SizeMb:          1024,
				PersistencyKind: computepb.VolumeRequest_CACHE,
			}},
			ExportPorts: []*computepb.ContainerPort{{
				Name:          "http",
				ContainerPort: 8080,
				Proto:         computepb.ContainerPort_HTTP,
			}},
		}},
	})
	if err != nil {
		return "", err
	}

	if _, err := cli.Compute.WaitInstanceSync(ctx, &computepb.WaitInstanceRequest{
		InstanceId: resp.Metadata.InstanceId,
	}); err != nil {
		return "", err
	}

	for _, container := range resp.Containers {
		if container.Name != "server" {
			continue
		}
		for _, port := range container.ExportedPort {
			if port.ContainerPort == 8080 && port.Proto == computepb.ContainerPort_HTTP {
				fmt.Fprintf(os.Stderr, "[namespace] Ensured instance: %s\n", resp.InstanceUrl)
				return port.Endpoint, nil
			}
		}
	}

	return "", fmt.Errorf("HTTP endpoint was not allocated")
}
