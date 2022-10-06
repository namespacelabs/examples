// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type putRequest struct {
	Key  string `json:"key"`
	Body []byte `json:"body"`
}

func put(ctx context.Context, cli *s3.Client) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed putRequest
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

type getRequest struct {
	Key string `json:"key"`
}

func get(ctx context.Context, cli *s3.Client) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		var parsed getRequest
		if err := json.NewDecoder(req.Body).Decode(&parsed); err != nil {
			rw.WriteHeader(400)
			fmt.Fprintf(rw, "invalid request: %v\n", err)
			return
		}

		if _, err := cli.GetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(parsed.Key),
		}); err != nil {
			rw.WriteHeader(500)
			fmt.Fprintf(rw, "failed to get object: %v\n", err)
			return
		}
	}
}
