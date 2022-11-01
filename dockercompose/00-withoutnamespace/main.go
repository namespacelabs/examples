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
		count := getHitCount(ctx, cache)
		fmt.Fprintf(rw, "Hello World! I have been seen %d times.\n", count)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	log.Printf("Listening on port: %d\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
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
