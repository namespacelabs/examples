server: {
	name: "strapi-frontend"

	integration: nodejs: {
		backends: strapibackend: "namespacelabs.dev/examples/strapi/02-with-postgres-s3/backend:backendapi"
	}

	services: webapi: {
		port: 3000
		kind: "http"

		ingress: internetFacing: true
	}
}
