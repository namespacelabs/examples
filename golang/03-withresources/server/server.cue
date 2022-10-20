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
			class:    "namespacelabs.dev/examples/golang/03-withresources/s3:Bucket"
			provider: "namespacelabs.dev/examples/golang/03-withresources/minio"

			intent: {
				region:     "us-east-1"
				bucketName: "testbucket"
			}
		}
	}
}

tests: {
	putAndGet: {
		builder: go: pkg: "./test"
	}
}
