// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "nextjsserver"

	integration: nodejs: {
		// As an optimization, you can specify a directory to copy files from in prod/test.
		// If not set, the whole app directory is copied.
		build: outDir: ".next"
	}

	// Through adding a resource here, Namespace will
	// 1) add Postgres server to the stack
	// 2) instantiate a Postgres database using Postgres server
	// 3) inject the configuration of the database (e.g. endpoint, password) into the resource config of our Go server
	resources: {
		myDatabase: {
			class:    "namespacelabs.dev/foundation/library/database/postgres:Database"
			provider: "namespacelabs.dev/foundation/library/oss/postgres"

			intent: {
				name: "nextjs"
				schema: ["schema.sql"]
			}

			resources: {
				// Select which cluster to host the Postgres database in. A foundation managed
				// colocated Postgres server is used for demonstration purposes.
				"cluster": "namespacelabs.dev/foundation/library/oss/postgres:colocated"
			}
		}
	}

	env: {
		POSTGRES_DB: fromResourceField: {
			resource: ":myDatabase"
			fieldRef: "name"
		}
		POSTGRES_ADDRESS: fromResourceField: {
			resource: ":myDatabase"
			fieldRef: "cluster_address"
		}
		POSTGRES_PASSWORD: fromResourceField: {
			resource: ":myDatabase"
			fieldRef: "password"
		}
	}

	services: webapi: {
		port: 3000
		kind: "http"

		ingress: true
	}
}
