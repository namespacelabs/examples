// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"namespacelabs.dev/foundation/schema/runtime"
	"namespacelabs.dev/foundation/std/go/core"
)

const (
	dbPackage   = "namespacelabs.dev/examples/multitier/01-simple/postgres"
	connBackoff = 500 * time.Millisecond
)

var (
	//go:embed schema.sql
	lib embed.FS
)

func main() {
	ctx := context.Background()
	config, err := core.LoadRuntimeConfig()
	if err != nil {
		panic(err)
	}

	conn, err := connectPG(ctx, config)
	if err != nil {
		panic(err)
	}

	schema, err := fs.ReadFile(lib, "schema.sql")
	if err != nil {
		panic(err)
	}
	if _, err := conn.Exec(ctx, string(schema)); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/add", add(ctx, conn))
	r.HandleFunc("/remove", remove(ctx, conn))
	r.HandleFunc("/list", list(ctx, conn))
	r.HandleFunc("/stream", stream(ctx, conn))

	r.HandleFunc("/readyz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "All OK\n\n")
	})

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r))
}

func connectPG(ctx context.Context, config *runtime.RuntimeConfig) (conn *pgx.Conn, err error) {
	db := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	var endpoint string
	for _, e := range config.StackEntry {
		if e.PackageName != dbPackage {
			continue
		}

		for _, s := range e.Service {
			if s.Name == "postgres" {
				endpoint = s.Endpoint
				break
			}
		}
	}

	cfg, err := pgx.ParseConfig(fmt.Sprintf("postgres://postgres:%s@%s/%s", password, endpoint, db))
	if err != nil {
		return nil, err
	}
	cfg.ConnectTimeout = connBackoff

	// Retry until backend is ready.
	err = backoff.Retry(func() error {
		conn, err = pgx.ConnectConfig(ctx, cfg)
		if err == nil {
			return nil
		}

		log.Printf("failed to connect to postgres: %v\n", err)
		return err
	}, backoff.WithContext(backoff.NewConstantBackOff(connBackoff), ctx))

	return conn, err
}
