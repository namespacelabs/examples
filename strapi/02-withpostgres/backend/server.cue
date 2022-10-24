server: {
	name: "strapi-backend"

	integration: "nodejs"

	env: {
		DATABASE_NAME: "bank"
		// Using a hard-coded passwords/keys to simplify this example.
		// See multitier/02-withsecrets/postgres example for how to use a generated secret as the password.
		DATABASE_PASSWORD: "DemoPasswordValue"
		APP_KEYS:          "testKey1,testKey2"
		API_TOKEN_SALT:    "testSalt"
		ADMIN_JWT_SECRET:  "testSecret"
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
		"namespacelabs.dev/examples/strapi/02-withpostgres/postgres",
	]
}
