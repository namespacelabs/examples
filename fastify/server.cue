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
