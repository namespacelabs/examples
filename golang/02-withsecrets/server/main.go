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
	"namespacelabs.dev/examples/shared"
	"namespacelabs.dev/foundation/framework/runtime"
	runtimepb "namespacelabs.dev/foundation/schema/runtime"
)

const (
	minioServer = "namespacelabs.dev/examples/golang/02-withsecrets/minio"
	bucketName  = "test-bucket"
	connBackoff = 500 * time.Millisecond
)

func main() {
	ctx := context.Background()
	config, err := runtime.LoadRuntimeConfig()
	if err != nil {
		log.Fatal(err)
	}

	cli, err := connectS3(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

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
	r.HandleFunc("/put", put(cli))
	r.HandleFunc("/get", get(cli))
	r.PathPrefix("/").HandlerFunc(shared.WelcomePage(config.Current))

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func connectS3(ctx context.Context, rtcfg *runtimepb.RuntimeConfig) (*s3.Client, error) {
	endpoint, err := runtime.ServerEndpoint(rtcfg, minioServer, "api")
	if err != nil {
		return nil, err
	}

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           fmt.Sprintf("http://%s", endpoint),
			SigningRegion: region,
		}, nil
	})

	region := os.Getenv("S3_REGION")
	accessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("S3_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
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
