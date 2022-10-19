resources: {
	minio: {
		kind: "namespacelabs.dev/examples/golang/03-withresources/s3:Bucket"
		on:   "namespacelabs.dev/examples/golang/03-withresources/minio"

		input: {
			region:     "us-east-1"
			bucketName: "testbucket"
		}
	}
}
