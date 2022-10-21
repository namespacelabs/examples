// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/framework/runtime"
	"namespacelabs.dev/foundation/library/storage/s3"
)

const dataBucketRef = "namespacelabs.dev/examples/golang/03-withresources/server:dataBucket"

func main() {
	ctx := context.Background()
	config, err := runtime.LoadRuntimeConfig()
	if err != nil {
		log.Fatal(err)
	}

	resources, err := resources.LoadResources()
	if err != nil {
		log.Fatal(err)
	}

	bucket := &s3.BucketInstance{}
	if err := resources.Unmarshal(dataBucketRef, bucket); err != nil {
		log.Fatal(err)
	}

	cli, err := connectS3(ctx, bucket)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/put", put(ctx, cli, bucket.BucketName))
	http.HandleFunc("/get", get(ctx, cli, bucket.BucketName))
	http.HandleFunc("/readyz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "All OK\n\n")
	})

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func connectS3(ctx context.Context, bucket *s3.BucketInstance) (*awss3.Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{PartitionID: "aws", URL: bucket.Url, SigningRegion: region}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(bucket.Region),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(bucket.AccessKey, bucket.SecretAccessKey, "" /* session */)))
	if err != nil {
		return nil, err
	}

	return awss3.NewFromConfig(cfg, func(o *awss3.Options) {
		o.UsePathStyle = true
	}), nil
}
