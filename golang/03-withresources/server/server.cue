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

	// TODO describe what this does
	resources: [
		"namespacelabs.dev/examples/golang/03-withresources/server/resources:minio",
	]
}

tests: {
	putAndGet: {
		builder: go: pkg: "./test"
	}
}
