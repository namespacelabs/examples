// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"

	"namespacelabs.dev/foundation/framework/pages"
	"namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/framework/runtime"
	"namespacelabs.dev/foundation/library/database/postgres"
)

const todosDatabaseRef = "namespacelabs.dev/examples/multitier/03-withresources/apibackend:todosDatabase"

func main() {
	ctx := context.Background()
	config, err := runtime.LoadRuntimeConfig()
	if err != nil {
		log.Fatal(err)
	}

	resources, err := resources.LoadResources()
	if err != nil {
		log.Fatal(err)
	}

	db := &postgres.DatabaseInstance{}
	if err := resources.Unmarshal(todosDatabaseRef, db); err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(ctx, db.ConnectionUri)
	if err != nil {
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
