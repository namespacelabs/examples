// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "go-server"

	integration: "go"

	env: {
		S3_BUCKET_NAME: "test-bucket"

		// Instructs Namespace to inject the secrets as environment variables to the container.
		// See multitier/02-withsecrets/postgres/server.cue for an example injecting secrets into a mount.
		S3_ACCESS_KEY_ID: fromSecret:     "namespacelabs.dev/examples/golang/02-withsecrets/minio:user"
		S3_SECRET_ACCESS_KEY: fromSecret: "namespacelabs.dev/examples/golang/02-withsecrets/minio:password"

		// Injects the endpoint of MinIO server into an environment variable.
		// Alternatively, could be read from /namespace/config/runtime.json.
		// See also https://github.com/namespacelabs/foundation/blob/main/framework/runtime/parsing.go
		S3_ENDPOINT: fromServiceEndpoint: "namespacelabs.dev/examples/golang/02-withsecrets/minio:api"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: true

			probe: http: "/readyz"
		}
	}

	// When adding a reference to Minio server to the `requires` block, Namespace will
	// 1) add Minio server to the deployed stack
	// 2) inject the configuration of Minio server (e.g. endpoint) into the runtime config of our Go server
	requires: [
		"namespacelabs.dev/examples/golang/02-withsecrets/minio",
	]
}

tests: {
	putAndGet: {
		integration: go: pkg: "./test"
	}
}
