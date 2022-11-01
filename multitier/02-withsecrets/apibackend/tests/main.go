// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"namespacelabs.dev/examples/multitier/02-withsecrets/apibackend/api"
	"namespacelabs.dev/foundation/framework/testing"
)

func main() {
	testing.Do(func(ctx context.Context, t testing.Test) error {
		e := t.MustEndpoint("namespacelabs.dev/examples/multitier/02-withsecrets/apibackend", "webapi")

		if err := t.WaitForEndpoint(ctx, e); err != nil {
			return err
		}

		items := []string{"item1", "item2"}

		for _, item := range items {
			if _, err := send(fmt.Sprintf("http://%s/add", e.Address()), &api.AddRequest{
				Name: item,
			}); err != nil {
				return err
			}
		}

		oldList, err := list(e.Address())
		if err != nil {
			return err
		}

		if len(oldList) != len(items) {
			return fmt.Errorf("expected %d items, got %d", len(items), len(oldList))
		}

		for i, item := range items {
			if oldList[i].Name != item {
				return fmt.Errorf("expected item %d to be %q but got %q", i, item, oldList[i].Name)
			}
		}

		if _, err := send(fmt.Sprintf("http://%s/remove", e.Address()), &api.RemoveRequest{
			Id: oldList[0].Id,
		}); err != nil {
			return err
		}

		newList, err := list(e.Address())
		if err != nil {
			return err
		}

		if len(newList) != len(oldList)-1 {
			return fmt.Errorf("expected %d items, got %d", len(oldList)-1, len(newList))
		}

		return nil
	})
}

func send(url string, v any) (*http.Response, error) {
	serialized, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(serialized))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s returned status code: %s", url, resp.Status)
	}

	return resp, nil
}

func list(endpoint string) ([]api.TodoItem, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/list", endpoint))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("list failed with status code: %s", resp.Status)
	}

	var decoded []api.TodoItem
	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return nil, err
	}

	return decoded, nil
}
