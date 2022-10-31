server: {
	name: "composetest"

	integration: "dockerfile"

	services: {
		web: {
			port: 5000
			kind: "http"

			ingress: internetFacing: true
		}
	}

	requires: [
		"namespacelabs.dev/foundation/universe/db/redis/server",
	]
}
