// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
	"namespacelabs.dev/foundation/framework/pages"
	"namespacelabs.dev/foundation/framework/runtime"
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

	// No need to create bucket anymore.

	r := mux.NewRouter()
	r.HandleFunc("/put", put(cli, bucketName))
	r.HandleFunc("/get", get(cli, bucketName))
	r.PathPrefix("/").HandlerFunc(pages.WelcomePage(config.Current))

	httpPort := os.Getenv("HTTP_PORT")
	log.Printf("Listening on port: %s\n", httpPort)
	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), r)
}

func connectS3(ctx context.Context) (*awss3.Client, error) {
	url := os.Getenv("S3_ENDPOINT_URL")

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{PartitionID: "aws", URL: url, SigningRegion: region}, nil
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

	return awss3.NewFromConfig(cfg, func(o *awss3.Options) {
		o.UsePathStyle = true
	}), nil
}
