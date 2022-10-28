// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
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

	// Through adding a resource here, Namespace will
	// 1) instatiate an S3 Bucket using MinIO
	// 2) inject the configuration of the bucket (e.g. endpoint, access keys) into the resource config of our Go server
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
		integration: go: pkg: "./test"
	}
}
