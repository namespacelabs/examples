// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cenkalti/backoff/v4"
	"namespacelabs.dev/examples/golang/03-withresources/s3"
	"namespacelabs.dev/foundation/schema/runtime"
)

const (
	server      = "namespacelabs.dev/examples/golang/03-withresources/minio:minioServer"
	connBackoff = 500 * time.Millisecond
)

var (
	intent    = flag.String("intent", "", "The serialized JSON intent.")
	resources = flag.String("resources", "", "The serialized JSON resources.")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	if *intent == "" {
		log.Fatal("--intent is missing")
	}

	i := &s3.S3Intent{}
	if err := json.Unmarshal([]byte(*intent), i); err != nil {
		log.Fatal(err)
	}

	endpoint, err := getEndpoint()
	if err != nil {
		log.Fatal(err)
	}

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           fmt.Sprintf("http://%s", endpoint),
			SigningRegion: region,
		}, nil
	})

	// TODO!!
	accessKeyID := os.Getenv("S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("S3_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(i.Region),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credProvider{accessKeyID: accessKeyID, secretAccessKey: secretAccessKey}))
	if err != nil {
		panic(err)
	}

	cli := awss3.NewFromConfig(cfg, func(o *awss3.Options) {
		o.UsePathStyle = true
	})

	// Retry until bucket is ready.
	log.Printf("Creating bucket %s.\n", i.BucketName)
	if err := backoff.Retry(func() error {
		// Speed up bucket creation through faster retries.
		ctx, cancel := context.WithTimeout(ctx, connBackoff)
		defer cancel()

		_, err := cli.CreateBucket(ctx, &awss3.CreateBucketInput{
			Bucket: aws.String(i.BucketName),
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
	log.Printf("Bucket %s created\n", i.BucketName)

	// TODO consume resources
	out := &s3.S3Instance{
		Region:          i.Region,
		AccessKey:       accessKeyID,
		SecretAccessKey: secretAccessKey,
		BucketName:      i.BucketName,
	}

	serialized, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("namespace.provision.result: %s\n", serialized)
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

func getEndpoint() (string, error) {
	r := make(map[string]any)
	if err := json.Unmarshal([]byte(*resources), &r); err != nil {
		return "", err
	}

	s, ok := r[server]
	if !ok {
		return "", fmt.Errorf("%s not found", server)
	}

	// TODO ahhhh
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	cfg := &runtime.Server{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return "", err
	}

	for _, s := range cfg.Service {
		if s.Name == "api" {
			return s.Endpoint, nil
		}
	}

	return "", fmt.Errorf("api endpoint not found")
}
