// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
server: {
	name: "go-backend-server"

	integration: "go"

	env: {
		POSTGRES_DB: "todos"

		// Using a hard-coded password to simplify this example.
		// See multitier/02-withsecrets/ for an example using generated secrets.
		POSTGRES_PASSWORD: "DemoPasswordValue"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: internetFacing: true

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
		imageFrom: shellscript: {
			entrypoint: "tests/curl.sh"
			requiredPackages: ["jq"]
		}
	}
	api: {
		imageFrom: go: pkg: "tests"
	}
}
