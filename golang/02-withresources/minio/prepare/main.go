// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"namespacelabs.dev/examples/golang/02-withresources/s3"
)

var intent = flag.String("intent", "", "The serialized JSON intent.")

func main() {
	_ = flag.String("resources", "", "The serialized JSON resources.")
	flag.Parse()

	if *intent == "" {
		log.Fatal("--intent is missing")
	}

	i := &s3.S3Intent{}
	if err := json.Unmarshal([]byte(*intent), i); err != nil {
		log.Fatal(err)
	}

	// TODO generate keys
	out := &s3.S3Instance{
		Region:          i.Region,
		AccessKey:       "TestOnlyUser",
		SecretAccessKey: "TestOnlyPassword",
	}

	// TODO add minio to stack

	serialized, err := json.Marshal(out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("namespace.provision.result: %s\n", serialized)
}
