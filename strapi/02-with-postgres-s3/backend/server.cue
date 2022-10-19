server: {
	name: "strapi-backend"

	integration: "nodejs"

	services: backendapi: {
		port: 1337
		kind: "http"

		// The data API endpoint needs to be publicly available for client-side rendering
		ingress: internetFacing: true
	}
}
