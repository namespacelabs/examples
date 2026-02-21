package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	computepb "buf.build/gen/go/namespace/cloud/protocolbuffers/go/proto/namespace/cloud/compute/v1beta"
	"buf.build/gen/go/namespace/cloud/grpc/go/proto/namespace/cloud/compute/v1beta/computev1betagrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"namespacelabs.dev/integrations/api"
	"namespacelabs.dev/integrations/api/compute"
	"namespacelabs.dev/integrations/auth"
	"namespacelabs.dev/integrations/nsc/grpcapi"
)

func main() {
	flag.Parse()

	token, err := auth.LoadDefaults()
	if err != nil {
		log.Fatal(err)
	}

	if err := run(context.Background(), os.Stderr, token, &computepb.InstanceShape{
		VirtualCpu:      2,
		MemoryMegabytes: 4 * 1024,
		MachineArch:     "amd64",
	}); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, debugLog io.Writer, token api.TokenSource, shape *computepb.InstanceShape) error {
	cli, err := compute.NewClient(ctx, token)
	if err != nil {
		return err
	}

	defer cli.Close()

	enc := json.NewEncoder(debugLog)
	enc.SetIndent("", "  ")

	start := time.Now()

	resp, err := cli.Compute.CreateInstance(ctx, &computepb.CreateInstanceRequest{
		Shape:             shape,
		DocumentedPurpose: "exec example",
		Deadline:          timestamppb.New(time.Now().Add(10 * time.Minute)),
		Containers: []*computepb.ContainerRequest{{
			Name:       "ubuntu",
			ImageRef:   "ubuntu:latest",
			Entrypoint: []string{"sleep", "600"},
			Args:       []string{},
		}},
	})
	if err != nil {
		return fmt.Errorf("failed to create instance: %w", err)
	}

	fmt.Fprintf(debugLog, "[namespace] Instance: %s\n", resp.InstanceUrl)
	_ = enc.Encode(resp)

	endpoint := resp.ExtendedMetadata.GetCommandServiceEndpoint()
	if endpoint == "" {
		return fmt.Errorf("command service endpoint not available")
	}

	fmt.Fprintf(debugLog, "[namespace] Command service endpoint: %s\n", endpoint)

	// Connect to the CommandService on the instance.
	conn, err := grpcapi.NewConnectionWithEndpoint(ctx, endpoint, token)
	if err != nil {
		return fmt.Errorf("failed to connect to command service: %w", err)
	}

	defer conn.Close()

	cmdCli := computev1betagrpc.NewCommandServiceClient(conn)

	result, err := cmdCli.RunCommandSync(ctx, &computepb.RunCommandRequest{
		InstanceId:          resp.Metadata.InstanceId,
		TargetContainerName: "ubuntu",
		Command: &computepb.Command{
			Command: []string{"uname", "-a"},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	fmt.Fprintf(debugLog, "[namespace] Total time from CreateInstance to command result: %s\n", time.Since(start))

	fmt.Fprintf(os.Stdout, "%s", result.Stdout)
	if len(result.Stderr) > 0 {
		fmt.Fprintf(os.Stderr, "%s", result.Stderr)
	}

	if result.ExitCode != 0 {
		return fmt.Errorf("command exited with code %d", result.ExitCode)
	}

	return nil
}
