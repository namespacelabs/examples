server: {
	name: "strapi-backend"

	integration: "nodejs"

	env: {
		DATABASE_NAME:     "bank"
		DATABASE_PASSWORD: "DemoPasswordValue"
	}

	services: backendapi: {
		port: 1337
		kind: "http"

		// The data API endpoint needs to be publicly available for client-side rendering
		ingress: internetFacing: true
	}

	// When adding a reference to Postgres server to the `requires` block, Namespace will
	// 1) add Postgres server to the deployed stack
	// 2) inject the configuration of Postgres server (e.g. endpoint) into the runtime config of Strapi server
	requires: [
		"namespacelabs.dev/examples/strapi/02-with-postgres-s3/postgres",
	]
}
