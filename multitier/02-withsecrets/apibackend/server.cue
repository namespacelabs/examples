server: {
	name: "go-backend-server"

	integration: "go"

	env: {
		POSTGRES_DB:            "todos"
		POSTGRES_PASSWORD_FILE: "/postgres/secrets/password"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: internetFacing: true

			probe: http: "/readyz"
		}
	}

	mounts: {
		"/postgres/secrets": configurable: {
			contents: {
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
	addAndList: {
		builder: shellscript: {
			entrypoint: "test/test.sh"
			requiredPackages: ["jq"]
		}
	}
}
