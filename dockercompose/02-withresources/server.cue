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

	// Through adding a resource here, Namespace will
	// 1) instatiate a cache using Redis
	// 2) inject the configuration of the cache (e.g. endpoint) into the resource config of our Python server
	resources: {
		cache: {
			// TODO!
			class:    "namespacelabs.dev/foundation/library/database:Database"
			provider: "namespacelabs.dev/foundation/library/oss/redis"
		}
	}
}
