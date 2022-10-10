server: {
	name: "nextjsserver"

	integration: "nodejs"

	services: webapi: {
		port: 3000
		kind: "http"

		ingress: internetFacing: true
	}

	// When adding a reference to Postgres server to the `requires` block, Namespace will
	// 1) add Postgres server to the deployed stack
	// 2) inject the configuration of Postgres server (e.g. endpoint) into the runtime config of our Next.js server
	requires: [
		"namespacelabs.dev/examples/nextjs/01-simple/postgres",
	]
}
