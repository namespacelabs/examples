server: {
	name: "go-server"

	integration: "go"

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: internetFacing: true

			probe: http: "/readyz"
		}
	}

	resources: {
		dataBucket: {
			class:    "namespacelabs.dev/foundation/library/storage/s3:Bucket"
			provider: "namespacelabs.dev/foundation/library/oss/minio"

			intent: {
				region:     "us-east-1"
				bucketName: "testbucket"
			}
		}
	}
}

tests: {
	putAndGet: {
		imageFrom: go: pkg: "./test"
	}
}
