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

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"namespacelabs.dev/foundation/framework/pages"
	"namespacelabs.dev/foundation/framework/runtime"
)

const httpPort = 4000 // Alternatively, could be read from /namespace/config/runtime.json.

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

	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://postgres:%s@%s/%s",
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("PG_ENDPOINT"), os.Getenv("POSTGRES_DB")))
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

	log.Printf("Listening on port: %d\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r))
}
