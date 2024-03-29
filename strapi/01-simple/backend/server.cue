// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "strapi-backend"

	integration: "nodejs"

	env: {
		// Using a hard-coded passwords/keys to simplify this example.
		// See multitier/02-withsecrets/postgres example for how to use generated secrets.
		APP_KEYS:         "testKey1,testKey2"
		API_TOKEN_SALT:   "testSalt"
		ADMIN_JWT_SECRET: "testSecret"
		JWT_SECRET:       "testSecret"
	}

	services: backendapi: {
		port: 1337
		kind: "http"

		// The data API endpoint needs to be publicly available for client-side rendering
		ingress: true
	}
}
