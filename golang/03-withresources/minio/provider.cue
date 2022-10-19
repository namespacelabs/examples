providers: {
	"namespacelabs.dev/examples/golang/03-withresources/s3:Bucket": {
		initializedWith: {
			binary: "namespacelabs.dev/examples/golang/03-withresources/minio/prepare"
		}

		resources: {
			// Adds the server to the stack
			minioServer: {
				kind: "namespacelabs.dev/foundation/std/runtime:Server"
				input: package_name: "namespacelabs.dev/examples/golang/03-withresources/minio/server"
			}
		}
	}
}
