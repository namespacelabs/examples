// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "nextjsserver"

	integration: "nodejs"

	env: {
		POSTGRES_DB: "nextjs"

		// Using a hard-coded password to simplify this example.
		// See nextjs/02-withsecrets/ for an example using managed secrets.
		POSTGRES_PASSWORD: "DemoPasswordValue"
	}

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
