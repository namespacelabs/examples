package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-redis/redis/v8"
	"namespacelabs.dev/foundation/framework/runtime"
)

const (
	retries     = 5
	connBackoff = 500 * time.Millisecond
)

func main() {
	ctx := context.Background()
	cache := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    os.Getenv("REDIS_URL"),
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("trigger\n")
		count := getHitCount(ctx, cache)
		fmt.Fprintf(rw, "Hello World! I have been seen %d times.\n", count)
	})

	config, err := runtime.LoadRuntimeConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start\n")

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getHitCount(ctx context.Context, cache *redis.Client) int {
	var count int

	// Retry until backend is ready.
	err := backoff.Retry(func() error {
		res := cache.Incr(ctx, "hits")
		count = int(res.Val())
		return res.Err()
	}, backoff.WithMaxRetries(backoff.NewConstantBackOff(connBackoff), retries))

	if err != nil {
		log.Fatalf("%v", err)
	}

	return count
}
