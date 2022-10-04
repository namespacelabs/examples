// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v4"

	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/runtime"
)

const schema = `CREATE TABLE IF NOT EXISTS list (
    Id INT GENERATED ALWAYS AS IDENTITY,
    Item varchar(255) NOT NULL,
    PRIMARY KEY(Id)
);`

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

	if _, err := conn.Exec(ctx, schema); err != nil {
		panic(err)
	}

	http.HandleFunc("/post", post(ctx, conn))
	http.HandleFunc("/list", list(ctx, conn))

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func connectPG(ctx context.Context, config *runtime.RuntimeConfig) (conn *pgx.Conn, err error) {
	password := os.Getenv("POSTGRES_PASSWORD")

	var endpoint string
	for _, e := range config.StackEntry {
		if e.PackageName != "namespacelabs.dev/examples/nextjs/db/postgres/server" {
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
		conn, err = pgx.Connect(ctx, fmt.Sprintf("postgres://postgres:%s@%s/postgres", password, endpoint))
		return err
	}, backoff.WithContext(backoff.NewConstantBackOff(time.Second), ctx))

	return conn, err
}

type request struct {
	Item string `json:"item"`
}

func post(ctx context.Context, conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed request
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		if _, err := conn.Exec(ctx, "INSERT INTO list (Item) VALUES ($1);", parsed.Item); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to insert into db: %v\n", err)
			return
		}
	}
}

func list(ctx context.Context, conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, _ *http.Request) {
		rows, err := conn.Query(ctx, "SELECT Item FROM list;")
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to connect to db: %v\n", err)
			return
		}
		defer rows.Close()

		var items []string
		for rows.Next() {
			var item string
			err = rows.Scan(&item)
			if err != nil {
				rw.WriteHeader(500)
				fmt.Fprintf(rw, "failed to process db response: %v\n", err)
				return
			}
			items = append(items, item)
		}

		for _, item := range items {
			fmt.Fprintln(rw, item)
		}
	}
}
