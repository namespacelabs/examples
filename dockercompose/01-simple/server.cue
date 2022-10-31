server: {
	name: "composetest"

	integration: "dockerfile"

	env: {
		FLASK_DEBUG: "True"
		REDIS_SERVICE: fromServiceEndpoint: "namespacelabs.dev/foundation/library/oss/redis/server:redis"
	}

	services: {
		web: {
			port: 5000
			kind: "http"

			ingress: internetFacing: true
		}
	}

	// TODO add sync mount

	requires: [
		"namespacelabs.dev/foundation/library/oss/redis/server",
	]
}
