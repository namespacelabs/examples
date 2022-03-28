package main

import (
	"context"
	"fmt"
	"sort"

	"namespacelabs.dev/examples/todo-app/api/todos"
	"namespacelabs.dev/foundation/testing"
)

func main() {
	testing.Do(func(ctx context.Context, t testing.Test) error {
		conn, err := t.Connect(ctx, t.MustEndpoint("namespacelabs.dev/examples/todo-app/api/todos", "todos"))
		if err != nil {
			return err
		}

		client := todos.NewTodosServiceClient(conn)

		tasks := []string{"homework", "landry"}

		for _, t := range tasks {
			req := todos.AddRequest{
				Name: t,
			}

			if _, err = client.Add(ctx, &req); err != nil {
				return err
			}
		}

		list, err := client.List(ctx, &todos.ListRequest{})
		if err != nil {
			return err
		}

		sort.Strings(tasks)
		sort.SliceStable(list.Items, func(i, j int) bool {
			return list.Items[i].Name < list.Items[j].Name
		})

		for i, t := range tasks {
			if t != list.Items[i].Name {
				return fmt.Errorf("item mismatch: '%v' is not '%v'", t, list.Items[i].Name)
			}
		}

		{
			req := todos.GetRelatedDataRequest{
				Id: list.Items[0].Id,
			}
			data, err := client.GetRelatedData(ctx, &req)
			if err != nil {
				return err
			}

			if data.Popularity < 1 || data.Popularity > 5 {
				return fmt.Errorf("invalid popularity score: '%v' is not in [1,5]", data.Popularity)
			}
		}

		req := todos.RemoveRequest{
			Id: list.Items[0].Id,
		}
		if _, err = client.Remove(ctx, &req); err != nil {
			return err
		}

		list, err = client.List(ctx, &todos.ListRequest{})
		if err != nil {
			return err
		}

		// "Testing" User Journey:
		// Uncomment next 3 lines once the bug with removing items has been fixed in "impl.go".
		// if len(list.Items) != len(tasks)-1 {
		// 	return fmt.Errorf("got %d items, expected %d", len(list.Items), len(tasks)-1)
		// }

		return nil
	})
}
