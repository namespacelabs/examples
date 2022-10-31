server: {
	name: "composetest"

	integration: "dockerfile"

	env: {
		FLASK_DEBUG: "True"
		// Injects the endpoint of Redis server into an environment variable.
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

	// When adding a reference to Redis server to the `requires` block, Namespace will
	// 1) add Redis server to the deployed stack
	// 2) inject the configuration of Redis server (e.g. endpoint) into the runtime config of our Python server.
	//    This enables referencing the configuration with fromServiceEndpoint above.
	requires: [
		"namespacelabs.dev/foundation/library/oss/redis/server",
	]
}
