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
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cenkalti/backoff/v4"
	"namespacelabs.dev/examples/golang/03-withresources/s3"
	fnresources "namespacelabs.dev/foundation/framework/resources"
	"namespacelabs.dev/foundation/schema/runtime"
	stdruntime "namespacelabs.dev/foundation/std/runtime"
)

const (
	providerPkg = "namespacelabs.dev/examples/golang/03-withresources/minio"
	connBackoff = 500 * time.Millisecond
)

var (
	intent    = flag.String("intent", "", "The serialized JSON intent.")
	resources = flag.String("resources", "", "The serialized JSON resources.")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	instance, err := createInstance()
	if err != nil {
		log.Fatal(err)
	}

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{PartitionID: "aws", URL: instance.Url, SigningRegion: region}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(instance.Region),
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(instance.AccessKey, instance.SecretAccessKey, "" /* session */)))
	if err != nil {
		log.Fatal(err)
	}

	cli := awss3.NewFromConfig(cfg, func(o *awss3.Options) {
		o.UsePathStyle = true
	})

	if err := createBucket(ctx, cli, instance.BucketName); err != nil {
		log.Fatal(err)
	}

	serialized, err := json.Marshal(instance)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("namespace.provision.result: %s\n", serialized)
}

func getEndpoint(resources *fnresources.Parser) (string, error) {
	key := fmt.Sprintf("%s:minioServer", providerPkg)
	cfg := &runtime.Server{}
	if err := resources.Decode(key, &cfg); err != nil {
		return "", err
	}

	for _, s := range cfg.Service {
		if s.Name == "api" {
			return s.Endpoint, nil
		}
	}

	return "", fmt.Errorf("api endpoint not found")
}

func readSecret(resources *fnresources.Parser, name string) (string, error) {
	key := fmt.Sprintf("%s:%s", providerPkg, name)
	secret := &stdruntime.SecretInstance{}
	if err := resources.Decode(key, &secret); err != nil {
		return "", err
	}

	data, err := os.ReadFile(secret.FilePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func createInstance() (*s3.BucketInstance, error) {
	if *intent == "" {
		return nil, fmt.Errorf("--intent is missing")
	}

	i := &s3.BucketIntent{}
	if err := json.Unmarshal([]byte(*intent), i); err != nil {
		return nil, err
	}

	if *resources == "" {
		return nil, fmt.Errorf("--resources is missing")
	}
	r := fnresources.NewParser([]byte(*resources))

	endpoint, err := getEndpoint(r)
	if err != nil {
		return nil, err
	}

	accessKeyID, err := readSecret(r, "minioUser")
	if err != nil {
		return nil, err
	}
	secretAccessKey, err := readSecret(r, "minioPassword")
	if err != nil {
		return nil, err
	}

	return &s3.BucketInstance{
		Region:          i.Region,
		AccessKey:       accessKeyID,
		SecretAccessKey: secretAccessKey,
		BucketName:      i.BucketName,
		Url:             fmt.Sprintf("http://%s", endpoint),
	}, nil
}

func createBucket(ctx context.Context, cli *awss3.Client, bucketName string) error {
	// Retry until bucket is ready.
	log.Printf("Creating bucket %s.\n", bucketName)
	if err := backoff.Retry(func() error {
		// Speed up bucket creation through faster retries.
		ctx, cancel := context.WithTimeout(ctx, connBackoff)
		defer cancel()

		_, err := cli.CreateBucket(ctx, &awss3.CreateBucketInput{
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
		return err
	}
	log.Printf("Bucket %s created\n", bucketName)

	return nil
}
