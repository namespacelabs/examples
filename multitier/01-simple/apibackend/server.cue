// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "go-backend-server"

	integration: "go"

	env: {
		POSTGRES_DB: "todos"

		// Using a hard-coded password to simplify this example.
		// See multitier/02-withsecrets/ for an example using generated secrets.
		POSTGRES_PASSWORD: "DemoPasswordValue"

		// Injects the endpoint of Postgres server into an environment variable.
		// Alternatively, could be read from /namespace/config/runtime.json.
		// See also https://github.com/namespacelabs/foundation/blob/main/framework/runtime/parsing.go
		PG_ENDPOINT: fromServiceEndpoint: "namespacelabs.dev/examples/multitier/01-simple/postgres:postgres"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: true

			probe: http: "/readyz"
		}
	}

	// When adding a reference to Postgres server to the `requires` block, Namespace will
	// 1) add Postgres server to the deployed stack
	// 2) inject the configuration of Postgres server (e.g. endpoint) into the runtime config of our Go backend server
	requires: [
		"namespacelabs.dev/examples/multitier/01-simple/postgres",
	]
}

tests: {
	simpleCurl: {
		integration: shellscript: "tests/curl.sh"
		env: {
			ENDPOINT: fromServiceEndpoint: ":webapi"
		}
	}

	api: {
		integration: go: pkg: "tests"
	}
}
