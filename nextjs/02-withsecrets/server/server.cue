// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
server: {
	name: "nextjsserver"

	integration: "nodejs"

	env: {
		POSTGRES_DB:            "nextjs"
		POSTGRES_PASSWORD_FILE: "/postgres/secrets/password"
	}

	services: webapi: {
		port: 3000
		kind: "http"

		ingress: internetFacing: true
	}

	mounts: {
		// Using an inline volume definition for brevity.
		// Mount points can also reference volumes by their package reference.
		// See multitier/02-withsecrets/postgres/server.cue for an example using a reference.
		"/postgres/secrets": configurable: {
			contents: {
				// Instructs Namespace to inject the secrets as file in the mount.
				// See golang/02-withsecrets/server/server.cue for an example injecting secrets as environment variables.
				"password": fromSecret: "namespacelabs.dev/examples/nextjs/02-withsecrets/postgres:password"
			}
		}
	}

	// When adding a reference to Postgres server to the `requires` block, Namespace will
	// 1) add Postgres server to the deployed stack
	// 2) inject the configuration of Postgres server (e.g. endpoint) into the runtime config of our Next.js server
	requires: [
		"namespacelabs.dev/examples/nextjs/02-withsecrets/postgres",
	]
}
