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
			kind: "namespacelabs.dev/examples/golang/03-withresources/s3:Bucket"
			on:   "namespacelabs.dev/examples/golang/03-withresources/minio"

			input: {
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
