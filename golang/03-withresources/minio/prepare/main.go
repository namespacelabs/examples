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

	if *intent == "" {
		log.Fatal("--intent is missing")
	}

	i := &s3.BucketIntent{}
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

	accessKeyID, err := readSecret("minioUser")
	if err != nil {
		log.Fatal(err)
	}
	secretAccessKey, err := readSecret("minioPassword")
	if err != nil {
		log.Fatal(err)
	}

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

		// TODO remove when we set the right secrets!
		if err != nil {
			panic(err)
		}

		log.Printf("failed to create bucket: %v\n", err)
		return err
	}, backoff.WithContext(backoff.NewConstantBackOff(connBackoff), ctx)); err != nil {
		panic(err)
	}
	log.Printf("Bucket %s created\n", i.BucketName)

	out := &s3.BucketInstance{
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
	key := fmt.Sprintf("%s:minioServer", providerPkg)
	cfg := &runtime.Server{}
	if err := fnresources.Decode([]byte(*resources), key, &cfg); err != nil {
		return "", err
	}

	for _, s := range cfg.Service {
		if s.Name == "api" {
			return s.Endpoint, nil
		}
	}

	return "", fmt.Errorf("api endpoint not found")
}

func readSecret(name string) (string, error) {
	key := fmt.Sprintf("%s:%s", providerPkg, name)
	secret := &stdruntime.SecretInstance{}
	if err := fnresources.Decode([]byte(*resources), key, &secret); err != nil {
		return "", err
	}

	data, err := os.ReadFile(secret.FilePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
