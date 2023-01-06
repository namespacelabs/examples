import "strings"

server: {
	name: "java-spring-boot-demo"

	// Using Dockerfile for a simple integration.
	// A native Java integration (on our roadmap) will allow best-in-class
	// build performance and minimal image size.
	integration: "dockerfile"

	// Through adding a resource here, Namespace will
	// 1) add Postgres server to the stack
	// 2) instantiate a Postgres database using Postgres server
	// 3) inject the configuration of the database (e.g. endpoint, password) into the resource config of our Java server
	resources: {
		todosDatabase: {
			class:    "namespacelabs.dev/foundation/library/database/postgres:Database"
			provider: "namespacelabs.dev/foundation/library/oss/postgres"

			intent: {
				name: "todos"
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
		// Injects the database instance fields into environment variables.
		SPRING_DATASOURCE_USERNAME: fromResourceField: {
			resource: ":todosDatabase"
			fieldRef: "user"
		}

		SPRING_DATASOURCE_PASSWORD: fromResourceField: {
			resource: ":todosDatabase"
			fieldRef: "password"
		}

		POSTGRES_ADDRESS: fromResourceField: {
			resource: ":todosDatabase"
			fieldRef: "cluster_address"
		}

		POSTGRES_DB: fromResourceField: {
			resource: ":todosDatabase"
			fieldRef: "name"
		}
	}

	args: [
		"java", "-jar", "/app/target/app.jar",
		"--spring.datasource.url=jdbc:postgresql://$(POSTGRES_ADDRESS)/$(POSTGRES_DB)",
	]

	services: {
		webapi: {
			port: 8080
			kind: "http"
		}
	}
}
