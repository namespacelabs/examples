// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"namespacelabs.dev/examples/multitier/03-withresources/apibackend/api"
	"namespacelabs.dev/go-ids"
)

func add(conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed api.AddRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		id := ids.NewSortableID()
		if _, err := conn.Exec(req.Context(), "INSERT INTO todos_table (Id, Name) VALUES ($1, $2);", id, parsed.Name); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to insert into db: %v\n", err)
			return
		}
	}
}

func remove(conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed api.RemoveRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		if _, err := conn.Exec(req.Context(), "DELETE FROM todos_table WHERE ID = $1;", parsed.Id); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to insert into db: %v\n", err)
			return
		}
	}
}

func fetchTodosList(ctx context.Context, conn *pgx.Conn) ([]api.TodoItem, error) {
	rows, err := conn.Query(ctx, "SELECT Id, Name FROM todos_table;")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	defer rows.Close()

	var todos []api.TodoItem
	for rows.Next() {
		var todo api.TodoItem
		err = rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to process db entry: %w", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func list(conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		todos, err := fetchTodosList(req.Context(), conn)
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "internal error: %v\n", err)
			return
		}

		serialized, err := json.Marshal(todos)
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "internal error: %v\n", err)
			return
		}
		fmt.Fprintln(rw, string(serialized))
	}
}

func stream(conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		// Poll the database for changes.
		t := time.NewTicker(250 * time.Millisecond)
		defer t.Stop()

		var previous []api.TodoItem

		for {
			select {
			case <-ctx.Done():
				return

			case <-t.C:
				todos, err := fetchTodosList(ctx, conn)
				if err != nil {
					rw.WriteHeader(500)
					fmt.Fprintf(rw, "internal error: %v\n", err)
					return
				}

				// Don't push data to the client, if it didn't change from the previous response.
				if equals(previous, todos) {
					continue
				}

				serialized, err := json.Marshal(todos)
				if err != nil {
					rw.WriteHeader(500)
					fmt.Fprintf(rw, "internal error: %v\n", err)
					return
				}
				fmt.Fprintln(rw, string(serialized))

				previous = todos
			}
		}
	}
}

func equals(a, b []api.TodoItem) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
