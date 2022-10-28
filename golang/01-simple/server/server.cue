// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
server: {
	name: "go-server"

	integration: "go"

	env: {
		S3_REGION: "us-east-1"

		// Using hard-coded credentials to simplify this example.
		// See golang/02-withsecrets/ for an example using managed secrets.
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

	// When adding a reference to Minio server to the `requires` block, Namespace will
	// 1) add Minio server to the deployed stack
	// 2) inject the configuration of Minio server (e.g. endpoint) into the runtime config of our Go server
	requires: [
		"namespacelabs.dev/examples/golang/01-simple/minio",
	]
}

tests: {
	putAndGet: {
		integration: go: pkg: "./test"
	}
}
