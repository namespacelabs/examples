server: {
	name: "strapi-frontend"

	integration: nodejs: {
		backends: strapibackend: "namespacelabs.dev/examples/strapi/01-simple/backend:backendapi"
	}

	services: webapi: {
		port: 3000
		kind: "http"

		ingress: internetFacing: true
	}
}
