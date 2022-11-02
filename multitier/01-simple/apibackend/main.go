// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

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
	"namespacelabs.dev/foundation/framework/pages"
	"namespacelabs.dev/foundation/framework/runtime"
	runtimepb "namespacelabs.dev/foundation/schema/runtime"
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
	config, err := runtime.LoadRuntimeConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := connectPG(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := fs.ReadFile(lib, "schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := conn.Exec(ctx, string(schema)); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/add", add(conn))
	r.HandleFunc("/remove", remove(conn))
	r.HandleFunc("/list", list(conn))
	r.HandleFunc("/stream", stream(conn))
	r.PathPrefix("/").HandlerFunc(pages.WelcomePage(config.Current))

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r))
}

func connectPG(ctx context.Context, config *runtimepb.RuntimeConfig) (conn *pgx.Conn, err error) {
	db := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	endpoint, err := runtime.ServerEndpoint(config, dbPackage, "postgres")
	if err != nil {
		return nil, err
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
