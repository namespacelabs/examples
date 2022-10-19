providers: {
	"namespacelabs.dev/examples/golang/03-withresources/s3:Bucket": {
		initializedWith: {
			binary: "namespacelabs.dev/examples/golang/03-withresources/minio/prepare"
		}

		// Adds the server to the stack
		resources: {
			minioServer: {
				kind: "namespacelabs.dev/foundation/std/runtime:Server"
				input: package_name: "namespacelabs.dev/examples/golang/03-withresources/minio/server"
			}
		}
	}
}
