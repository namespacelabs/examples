server: {
	name: "go-backend-server"

	integration: "go"

	env: {
		POSTGRES_DB:       "todos"
		POSTGRES_PASSWORD: "DemoPasswordValue"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: internetFacing: true
		}
	}
}

// TODO move requires into server when nsboot supports new syntax.
//
// When adding a reference to Postgres server to the `requires` block, Namespace will
// 1) add Postgres server to the deployed stack
// 2) inject the configuration of Postgres server (e.g. endpoint) into the runtime config of our Go backend server
requires: [
	"namespacelabs.dev/examples/multitier/01-simple/postgres",
]
