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
	"github.com/jackc/pgx/v4"

	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/runtime"
)

const dbPackage = "namespacelabs.dev/examples/multitier/simple/db"

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

	http.HandleFunc("/add", add(ctx, conn))
	http.HandleFunc("/remove", remove(ctx, conn))
	http.HandleFunc("/list", list(ctx, conn))
	http.HandleFunc("/stream", stream(ctx, conn))

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
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

	// Retry until backend is ready.
	err = backoff.Retry(func() error {
		conn, err = pgx.Connect(ctx, fmt.Sprintf("postgres://postgres:%s@%s/%s", password, endpoint, db))
		return err
	}, backoff.WithContext(backoff.NewConstantBackOff(time.Second), ctx))

	return conn, err
}
