// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
server: {
	name: "fastifyserver"

	integration: "nodejs"

	services: webapi: {
		port: 4000
		kind: "http"

		ingress: internetFacing: true
	}
}

tests: {
	smoke: {
		integration: nodejs: pkg: "test"
	}
}
