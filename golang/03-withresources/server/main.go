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
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"namespacelabs.dev/examples/golang/03-withresources/s3"
	fnresources "namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/std/go/core"
)

const minioResource = "namespacelabs.dev/examples/golang/03-withresources/server/resources:minio"

func main() {
	ctx := context.Background()
	config, err := core.LoadRuntimeConfig()
	if err != nil {
		panic(err)
	}

	resources, err := core.LoadResources()
	if err != nil {
		panic(err)
	}

	cli, err := connectS3(ctx, resources)
	if err != nil {
		panic(err)
	}

	bucketname := "TODO from resource"

	http.HandleFunc("/put", put(ctx, cli, bucketname))
	http.HandleFunc("/get", get(ctx, cli, bucketname))
	http.HandleFunc("/readyz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "All OK\n\n")
	})

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func connectS3(ctx context.Context, resources *fnresources.Parser) (*awss3.Client, error) {
	var bucket s3.BucketInstance

	if err := resources.Decode(minioResource, &bucket); err != nil {
		return nil, err
	}

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{PartitionID: "aws", URL: bucket.Url, SigningRegion: region}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(bucket.Region),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credProvider{accessKeyID: bucket.AccessKey, secretAccessKey: bucket.SecretAccessKey}))
	if err != nil {
		return nil, err
	}

	return awss3.NewFromConfig(cfg, func(o *awss3.Options) {
		o.UsePathStyle = true
	}), nil
}

type credProvider struct {
	accessKeyID     string
	secretAccessKey string
}

var _ aws.CredentialsProvider = credProvider{}

func (cf credProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     cf.accessKeyID,
		SecretAccessKey: cf.secretAccessKey,
	}, nil
}
