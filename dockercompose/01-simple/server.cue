server: {
	name: "composetest"

	integration: "dockerfile"

	env: {
		FLASK_DEBUG:   "True"
		REDIS_SERVICE: "redis-h1ic17p47gr4df9i" // TODO inject
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
		// TODO use lib instead.
		"namespacelabs.dev/foundation/universe/db/redis/server",
	]
}
