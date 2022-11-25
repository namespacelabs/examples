// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "go-backend-server"

	integration: "go"

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: true

			probe: http: "/readyz"
		}
	}

	// Through adding a resource here, Namespace will
	// 1) add Postgres server to the stack
	// 2) instatiate a Postgres database using Postgres server
	// 3) inject the configuration of the database (e.g. endpoint, password) into the resource config of our Go server
	resources: {
		todosDatabase: {
			class:    "namespacelabs.dev/foundation/library/database/postgres:Database"
			provider: "namespacelabs.dev/foundation/library/oss/postgres"

			intent: {
				name: "todos"
				schema: ["schema.sql"]
			}

			resources: {
				// Select which cluster to host the Postgres database in.
				// We use a local package reference to refer to the resource below.
				"cluster": ":postgresCluster"
			}
		}
	}
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

resources: {
	postgresCluster: {
		class:    "namespacelabs.dev/foundation/library/database/postgres:Cluster"
		provider: "namespacelabs.dev/foundation/library/oss/postgres"

		intent: {}
	}
}
