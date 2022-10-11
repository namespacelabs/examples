server: {
	name: "go-server"

	integration: "go"

	env: {
		S3_REGION:            "us-east-1"
		S3_ACCESS_KEY_ID:     "TestOnlyUser"
		S3_SECRET_ACCESS_KEY: "TestOnlyPassword"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: internetFacing: true

			probe: http: "/readyz"
		}
	}

	// When adding a reference to S3 server to the `requires` block, Namespace will
	// 1) add S3 server to the deployed stack
	// 2) inject the configuration of S3 server (e.g. endpoint) into the runtime config of our Go server
	requires: [
		"namespacelabs.dev/examples/golang/01-simple/s3",
	]
}

tests: {
	putAndGet: {
		build: go: pkg: "./test"
	}
}
