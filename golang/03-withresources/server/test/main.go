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

	"namespacelabs.dev/examples/golang/03-withresources/server/api"
	"namespacelabs.dev/foundation/framework/testing"
)

const (
	fakeKey     = "FooBar"
	fakeContent = "fake content"
)

func main() {
	testing.Do(func(ctx context.Context, t testing.Test) error {
		e := t.MustEndpoint("namespacelabs.dev/examples/golang/03-withresources/server", "webapi")

		if err := t.WaitForEndpoint(ctx, e); err != nil {
			return err
		}

		if _, err := send(fmt.Sprintf("http://%s/put", e.Address()), &api.PutRequest{
			Key:  fakeKey,
			Body: []byte(fakeContent),
		}); err != nil {
			return err
		}

		resp, err := send(fmt.Sprintf("http://%s/get", e.Address()), &api.GetRequest{
			Key: fakeKey,
		})
		if err != nil {
			return err
		}

		var getResp api.GetResponse
		if err := json.NewDecoder(resp.Body).Decode(&getResp); err != nil {
			return err
		}

		if string(getResp.Body) != fakeContent {
			return fmt.Errorf("expected content %q but got %q", fakeContent, string(getResp.Body))
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
