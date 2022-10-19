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
			// Mounts the minio user as a secret
			minioUser: {
				kind: "namespacelabs.dev/foundation/std/runtime:Secret"
				input: ref: "namespacelabs.dev/examples/golang/03-withresources/minio/server:user"
			}
			// Mounts the minio password as a secret
			minioPassword: {
				kind: "namespacelabs.dev/foundation/std/runtime:Secret"
				input: ref: "namespacelabs.dev/examples/golang/03-withresources/minio/server:password"
			}
		}
	}
}
