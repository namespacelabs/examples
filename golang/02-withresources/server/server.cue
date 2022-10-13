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

	// TODO describe what this does
	resources: [
		"namespacelabs.dev/examples/golang/02-withresources/server/resources:minio",
	]
}

tests: {
	putAndGet: {
		builder: go: pkg: "./test"
	}
}
