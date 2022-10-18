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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"namespacelabs.dev/examples/golang/01-simple/server/api"
)

func put(ctx context.Context, cli *s3.Client) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed api.PutRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		if _, err := cli.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(parsed.Key),
			Body:   bytes.NewReader(parsed.Body),
		}); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to upload object: %v\n", err)
			return
		}
	}
}

func get(ctx context.Context, cli *s3.Client) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed api.GetRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		out, err := cli.GetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(parsed.Key),
		})
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to get object: %v\n", err)
			return
		}

		content, err := io.ReadAll(out.Body)
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to read object: %v\n", err)
			return
		}

		resp := api.GetResponse{
			Body: content,
		}

		serialized, err := json.Marshal(resp)
		if err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "internal error: %v\n", err)
			return
		}

		fmt.Fprintln(rw, string(serialized))
	}
}
