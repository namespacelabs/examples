// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "go-backend-server"

	integration: "go"

	env: {
		POSTGRES_DB:            "todos"
		POSTGRES_PASSWORD_FILE: "/postgres/secrets/password"

		// Injects the endpoint of Postgres server into an environment variable.
		// Alternatively, could be read from /namespace/config/runtime.json.
		// See also https://github.com/namespacelabs/foundation/blob/main/framework/runtime/parsing.go
		PG_ENDPOINT: fromServiceEndpoint: "namespacelabs.dev/examples/multitier/02-withsecrets/postgres:postgres"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: true

			probe: http: "/readyz"
		}
	}

	mounts: {
		// Using an inline volume definition for brevity.
		// Mount points can also reference volumes by their package reference.
		// See multitier/02-withsecrets/postgres/server.cue for an example using a reference.
		"/postgres/secrets": configurable: {
			contents: {
				// Instructs Namespace to inject the secrets as file in the mount.
				// See golang/02-withsecrets/server/server.cue for an example injecting secrets as environment variables.
				"password": fromSecret: "namespacelabs.dev/examples/multitier/02-withsecrets/postgres:password"
			}
		}
	}

	// When adding a reference to Postgres server to the `requires` block, Namespace will
	// 1) add Postgres server to the deployed stack
	// 2) inject the configuration of Postgres server (e.g. endpoint) into the runtime config of our Go backend server
	requires: [
		"namespacelabs.dev/examples/multitier/02-withsecrets/postgres",
	]
}

tests: {
	simpleCurl: {
		integration: shellscript: entrypoint: "tests/curl.sh"
		env: {
			ENDPOINT: fromServiceEndpoint: ":webapi"
		}
	}

	api: {
		integration: go: pkg: "tests"
	}
}
