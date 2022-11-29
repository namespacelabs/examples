// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/mux"
	"namespacelabs.dev/foundation/framework/pages"
	"namespacelabs.dev/foundation/framework/runtime"
)

const (
	connBackoff = 500 * time.Millisecond
	httpPort    = 4000 // Alternatively, could be read from /namespace/config/runtime.json.
)

func main() {
	ctx := context.Background()
	config, err := runtime.LoadRuntimeConfig()
	if err != nil {
		log.Fatal(err)
	}

	cli, err := connectS3(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bucketName := os.Getenv("S3_BUCKET_NAME")

	// Retry until bucket is ready.
	log.Printf("Creating bucket %s.\n", bucketName)
	if err = backoff.Retry(func() error {
		// Speed up bucket creation through faster retries.
		ctx, cancel := context.WithTimeout(ctx, connBackoff)
		defer cancel()

		_, err := cli.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		var alreadyExists *types.BucketAlreadyExists
		var alreadyOwned *types.BucketAlreadyOwnedByYou
		if err == nil || errors.As(err, &alreadyExists) || errors.As(err, &alreadyOwned) {
			return nil
		}

		log.Printf("failed to create bucket: %v\n", err)
		return err
	}, backoff.WithContext(backoff.NewConstantBackOff(connBackoff), ctx)); err != nil {
		log.Fatal(err)
	}
	log.Printf("Bucket %s created\n", bucketName)

	r := mux.NewRouter()
	r.HandleFunc("/put", put(cli, bucketName))
	r.HandleFunc("/get", get(cli, bucketName))
	r.PathPrefix("/").HandlerFunc(pages.WelcomePage(config.Current))

	log.Printf("Listening on port: %d\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
}

func connectS3(ctx context.Context) (*s3.Client, error) {
	endpoint := os.Getenv("S3_ENDPOINT")

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           fmt.Sprintf("http://%s", endpoint),
			SigningRegion: region,
		}, nil
	})

	accessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("S3_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "" /* session */)))
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	}), nil
}
