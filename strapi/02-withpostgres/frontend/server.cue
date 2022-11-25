// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "strapi-frontend"

	integration: nodejs: {
		// When adding a reference to the Strapi backend to the `backends` block, Namespace will
		// 1) add the backend server to the deployed stack
		// 2) inject the configuration of the backend server (e.g. endpoint) into a src/config/backends.ns.js file that is accessible to the browser
		backends: strapibackend: "namespacelabs.dev/examples/strapi/02-withpostgres/backend:backendapi"
	}

	services: webapi: {
		port: 3000
		kind: "http"

		ingress: true
	}
}
