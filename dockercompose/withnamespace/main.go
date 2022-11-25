package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
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

	log.Printf("Listening on port: %s\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)
}

func getHitCount(ctx context.Context, cache *redis.Client) int {
	// No need to retry. Namespace ensures that Redis is ready.
	res := cache.Incr(ctx, "hits")

	if res.Err() != nil {
		log.Fatalf("%v", res.Err())
	}

	return int(res.Val())
}
