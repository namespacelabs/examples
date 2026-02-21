package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"

	computepb "buf.build/gen/go/namespace/cloud/protocolbuffers/go/proto/namespace/cloud/compute/v1beta"
	"buf.build/gen/go/namespace/cloud/grpc/go/proto/namespace/cloud/compute/v1beta/computev1betagrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"namespacelabs.dev/integrations/api"
	"namespacelabs.dev/integrations/api/compute"
	"namespacelabs.dev/integrations/auth"
	"namespacelabs.dev/integrations/nsc/grpcapi"
)

var (
	count = flag.Int("count", 100, "Number of instances to create in parallel.")
)

type result struct {
	Index      int
	InstanceID string
	CreateDur  time.Duration
	RunDur     time.Duration
	TotalDur   time.Duration
	Err        error
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	token, err := auth.LoadDefaults()
	if err != nil {
		log.Fatal(err)
	}

	cli, err := compute.NewClient(ctx, token)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	var (
		mu          sync.Mutex
		results     []result
		instanceIDs []string
		wg          sync.WaitGroup
	)

	fmt.Fprintf(os.Stderr, "Launching %d instances in parallel...\n", *count)
	start := time.Now()

	for i := 0; i < *count; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			r := runOne(ctx, cli, token, idx)
			mu.Lock()
			results = append(results, r)
			if r.InstanceID != "" {
				instanceIDs = append(instanceIDs, r.InstanceID)
			}
			mu.Unlock()
			if r.Err != nil {
				fmt.Fprintf(os.Stderr, "[%3d] FAIL %s: %v\n", idx, r.InstanceID, r.Err)
			} else {
				fmt.Fprintf(os.Stderr, "[%3d] OK   %s  create=%s run=%s total=%s\n",
					idx, r.InstanceID, r.CreateDur.Round(time.Millisecond), r.RunDur.Round(time.Millisecond), r.TotalDur.Round(time.Millisecond))
			}
		}(i)
	}

	wg.Wait()
	wallTime := time.Since(start)

	printStats(results, wallTime)

	// Clean up: destroy all instances.
	fmt.Fprintf(os.Stderr, "\nDestroying %d instances...\n", len(instanceIDs))
	var destroyWg sync.WaitGroup
	for _, id := range instanceIDs {
		destroyWg.Add(1)
		go func(id string) {
			defer destroyWg.Done()
			_, err := cli.Compute.DestroyInstance(context.Background(), &computepb.DestroyInstanceRequest{
				InstanceId: id,
				Reason:     "benchmark complete",
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "  destroy %s: %v\n", id, err)
			}
		}(id)
	}
	destroyWg.Wait()
	fmt.Fprintf(os.Stderr, "Done.\n")
}

func runOne(ctx context.Context, cli compute.Client, token api.TokenSource, idx int) result {
	r := result{Index: idx}
	totalStart := time.Now()

	// Create instance.
	createStart := time.Now()
	resp, err := cli.Compute.CreateInstance(ctx, &computepb.CreateInstanceRequest{
		Shape: &computepb.InstanceShape{
			VirtualCpu:      1,
			MemoryMegabytes: 2 * 1024,
			MachineArch:     "amd64",
		},
		DocumentedPurpose: "benchmark",
		Deadline:          timestamppb.New(time.Now().Add(1 * time.Hour)),
		Containers: []*computepb.ContainerRequest{{
			Name:       "ubuntu",
			ImageRef:   "ubuntu@sha256:d1e2e92c075e5ca139d51a140fff46f84315c0fdce203eab2807c7e495eff4f9",
			Entrypoint: []string{"sleep", "3600"},
			Args:       []string{},
		}},
		Experimental: &computepb.CreateInstanceRequest_ExperimentalFeatures{
			PrivateFeature: []string{"EXP_CONTAINER_USE_ROOT_DISK"},
		},
	})
	if err != nil {
		r.Err = fmt.Errorf("create: %w", err)
		return r
	}
	r.CreateDur = time.Since(createStart)
	r.InstanceID = resp.Metadata.InstanceId

	// Run command.
	endpoint := resp.ExtendedMetadata.GetCommandServiceEndpoint()
	if endpoint == "" {
		r.Err = fmt.Errorf("no command service endpoint")
		return r
	}

	conn, err := grpcapi.NewConnectionWithEndpoint(ctx, endpoint, token)
	if err != nil {
		r.Err = fmt.Errorf("connect: %w", err)
		return r
	}
	defer conn.Close()

	cmdCli := computev1betagrpc.NewCommandServiceClient(conn)

	runStart := time.Now()
	res, err := cmdCli.RunCommandSync(ctx, &computepb.RunCommandRequest{
		InstanceId:          resp.Metadata.InstanceId,
		TargetContainerName: "ubuntu",
		Command: &computepb.Command{
			Command: []string{"uname", "-a"},
		},
	})
	if err != nil {
		r.Err = fmt.Errorf("run: %w", err)
		return r
	}
	r.RunDur = time.Since(runStart)

	if res.ExitCode != 0 {
		r.Err = fmt.Errorf("exit code %d", res.ExitCode)
		return r
	}

	r.TotalDur = time.Since(totalStart)
	return r
}

func printStats(results []result, wallTime time.Duration) {
	var success, fail int
	var createDurs, runDurs, totalDurs []float64

	for _, r := range results {
		if r.Err != nil {
			fail++
			continue
		}
		success++
		createDurs = append(createDurs, r.CreateDur.Seconds())
		runDurs = append(runDurs, r.RunDur.Seconds())
		totalDurs = append(totalDurs, r.TotalDur.Seconds())
	}

	fmt.Fprintf(os.Stdout, "\n=== Benchmark Results ===\n")
	fmt.Fprintf(os.Stdout, "Instances: %d total, %d success, %d failed\n", len(results), success, fail)
	fmt.Fprintf(os.Stdout, "Wall time: %s\n\n", wallTime.Round(time.Millisecond))

	if success == 0 {
		return
	}

	printDurStats(os.Stdout, "Create Instance", createDurs)
	printDurStats(os.Stdout, "Run Command", runDurs)
	printDurStats(os.Stdout, "Total (create+run)", totalDurs)

	// Sort successful results by total duration.
	var successful []result
	for _, r := range results {
		if r.Err == nil {
			successful = append(successful, r)
		}
	}
	sort.Slice(successful, func(i, j int) bool {
		return successful[i].TotalDur < successful[j].TotalDur
	})

	n := 5
	if len(successful) < n {
		n = len(successful)
	}

	fmt.Fprintf(os.Stdout, "5 Fastest:\n")
	for i := 0; i < n; i++ {
		r := successful[i]
		fmt.Fprintf(os.Stdout, "  %s  total=%s  create=%s  run=%s\n",
			r.InstanceID, r.TotalDur.Round(time.Millisecond), r.CreateDur.Round(time.Millisecond), r.RunDur.Round(time.Millisecond))
	}

	fmt.Fprintf(os.Stdout, "\n5 Slowest:\n")
	for i := len(successful) - 1; i >= len(successful)-n; i-- {
		r := successful[i]
		fmt.Fprintf(os.Stdout, "  %s  total=%s  create=%s  run=%s\n",
			r.InstanceID, r.TotalDur.Round(time.Millisecond), r.CreateDur.Round(time.Millisecond), r.RunDur.Round(time.Millisecond))
	}
	fmt.Fprintf(os.Stdout, "\n")
}

func printDurStats(w *os.File, label string, vals []float64) {
	sort.Float64s(vals)
	n := len(vals)

	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	avg := sum / float64(n)

	variance := 0.0
	for _, v := range vals {
		variance += (v - avg) * (v - avg)
	}
	stddev := math.Sqrt(variance / float64(n))

	fmt.Fprintf(w, "%s:\n", label)
	fmt.Fprintf(w, "  min=%7.3fs  max=%7.3fs  avg=%7.3fs  stddev=%7.3fs\n", vals[0], vals[n-1], avg, stddev)
	fmt.Fprintf(w, "  p50=%7.3fs  p90=%7.3fs  p95=%7.3fs  p99=%7.3fs\n",
		percentile(vals, 50), percentile(vals, 90), percentile(vals, 95), percentile(vals, 99))
	fmt.Fprintf(w, "\n")
}

func percentile(sorted []float64, p float64) float64 {
	idx := (p / 100.0) * float64(len(sorted)-1)
	lower := int(idx)
	upper := lower + 1
	if upper >= len(sorted) {
		return sorted[len(sorted)-1]
	}
	frac := idx - float64(lower)
	return sorted[lower]*(1-frac) + sorted[upper]*frac
}
