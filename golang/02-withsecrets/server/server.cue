server: {
	name: "go-server"

	integration: "go"

	env: {
		S3_REGION: "us-east-1"

		// Instructs Namespace to inject the secrets as environment variables to the container.
		// See multitier/02-withsecrets/postgres/server.cue for an example injecting secrets into a mount.
		S3_ACCESS_KEY_ID: fromSecret:     "namespacelabs.dev/examples/golang/02-withsecrets/minio:user"
		S3_SECRET_ACCESS_KEY: fromSecret: "namespacelabs.dev/examples/golang/02-withsecrets/minio:password"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: internetFacing: true

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
		builder: go: pkg: "./test"
	}
}
