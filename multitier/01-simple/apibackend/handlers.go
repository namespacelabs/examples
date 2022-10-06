// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
	"namespacelabs.dev/go-ids"
)

type addRequest struct {
	Name string `json:"name"`
}

func add(ctx context.Context, conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed addRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		id := ids.NewSortableID()
		if _, err := conn.Exec(ctx, "INSERT INTO todos_table (Id, Name) VALUES ($1, $2);", id, parsed.Name); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to insert into db: %v\n", err)
			return
		}
	}
}

type removeRequest struct {
	Id string `json:"id"`
}

func remove(ctx context.Context, conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed removeRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		if _, err := conn.Exec(ctx, "DELETE FROM todos_table WHERE ID = $1;", parsed.Id); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to insert into db: %v\n", err)
			return
		}
	}
}

type todoItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func fetchTodosList(ctx context.Context, conn *pgx.Conn) ([]todoItem, error) {
	rows, err := conn.Query(ctx, "SELECT Id, Name FROM todos_table;")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	defer rows.Close()

	var todos []todoItem
	for rows.Next() {
		var todo todoItem
		err = rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to process db entry: %w", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func list(ctx context.Context, conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, _ *http.Request) {
		todos, err := fetchTodosList(ctx, conn)
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "internal error: %v\n", err)
			return
		}

		for _, todo := range todos {
			fmt.Fprintln(rw, todo)
		}
	}
}

func stream(ctx context.Context, conn *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, _ *http.Request) {
		// Poll the database for changes.
		t := time.NewTicker(250 * time.Millisecond)
		defer t.Stop()

		var previous []todoItem

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

				for _, item := range todos {
					fmt.Fprintln(rw, item)
				}

				previous = todos
			}
		}
	}
}

func equals(a, b []todoItem) bool {
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
