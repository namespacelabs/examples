// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

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
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cenkalti/backoff/v4"
	"namespacelabs.dev/foundation/schema/runtime"
	"namespacelabs.dev/foundation/std/go/core"
)

const (
	s3Package   = "namespacelabs.dev/examples/golang/02-withsecrets/s3"
	bucketName  = "test-bucket"
	connBackoff = 500 * time.Millisecond
)

func main() {
	ctx := context.Background()
	config, err := core.LoadRuntimeConfig()
	if err != nil {
		panic(err)
	}

	cli, err := connectS3(ctx, config)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	log.Printf("Bucket %s created\n", bucketName)

	http.HandleFunc("/put", put(ctx, cli))
	http.HandleFunc("/get", get(ctx, cli))
	http.HandleFunc("/readyz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "All OK\n\n")
	})

	port := config.Current.Port[0].Port
	log.Printf("Listening on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func connectS3(ctx context.Context, rtcfg *runtime.RuntimeConfig) (*s3.Client, error) {
	var endpoint string
	for _, e := range rtcfg.StackEntry {
		if e.PackageName != s3Package {
			continue
		}

		for _, s := range e.Service {
			if s.Name == "api" {
				endpoint = s.Endpoint
				break
			}
		}
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
			credProvider{accessKeyID: accessKeyID, secretAccessKey: secretAccessKey}))
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
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
