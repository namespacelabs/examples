providers: {
	"namespacelabs.dev/examples/golang/03-withresources/s3:Bucket": {
		initializedWith: {
			binary: "namespacelabs.dev/examples/golang/03-withresources/minio/prepare"
		}

		resources: {
			// Adds the server to the stack
			minioServer: {
				class: "namespacelabs.dev/foundation/std/runtime:Server"
				intent: package_name: "namespacelabs.dev/examples/golang/03-withresources/minio/server"
			}
			// Mounts the minio user as a secret
			minioUser: {
				class: "namespacelabs.dev/foundation/std/runtime:Secret"
				intent: ref: "namespacelabs.dev/examples/golang/03-withresources/minio/server:user"
			}
			// Mounts the minio password as a secret
			minioPassword: {
				class: "namespacelabs.dev/foundation/std/runtime:Secret"
				intent: ref: "namespacelabs.dev/examples/golang/03-withresources/minio/server:password"
			}
		}
	}
}
