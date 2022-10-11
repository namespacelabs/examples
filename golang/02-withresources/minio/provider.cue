providers: {
	"namespacelabs.dev/examples/golang/02-withresources/s3:S3": {
		initializedWith: {
			binary: "namespacelabs.dev/examples/golang/02-withresources/minio/prepare"
		}

		// Adds the server to the stack
		resources: {
			minioServer: {
				kind: "namespacelabs.dev/foundation/std/runtime:Server"
				input: package_name: "namespacelabs.dev/examples/golang/02-withresources/minio/server"
			}
		}
	}
}
