// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
		e := t.MustEndpoint("namespacelabs.dev/examples/golang/02-withsecrets/server", "webapi")

		if err := t.WaitForEndpoint(ctx, e); err != nil {
			return err
		}

		resp, err := send(fmt.Sprintf("http://%s/put", e.Address()), &api.PutRequest{
			Key:  fakeKey,
			Body: []byte(fakeContent),
		})
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			x, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return fmt.Errorf("put failed with status code: %s\n%s", resp.Status, string(x))
		}

		resp, err = send(fmt.Sprintf("http://%s/get", e.Address()), &api.GetRequest{
			Key: fakeKey,
		})
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("get failed with status code: %s", resp.Status)
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

	return http.Post(url, "application/json", bytes.NewReader(serialized))
}
