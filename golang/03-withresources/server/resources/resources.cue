resources: {
	minio: {
		kind: "namespacelabs.dev/examples/golang/03-withresources/s3:S3"
		on:   "namespacelabs.dev/examples/golang/03-withresources/minio"

		input: {
			region: "us-east-1"
		}
	}
}
