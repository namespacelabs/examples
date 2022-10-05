server: {
	name: "go-backend-server"

	integration: "go"

	env: {
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

requires: [
	"namespacelabs.dev/examples/multitier/simple/db",
]
