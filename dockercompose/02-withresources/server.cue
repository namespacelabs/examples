server: {
	name: "composetest"

	integration: "dockerfile"

	env: {
		// TODO add when we have code sync
		// FLASK_DEBUG: "True"

		REDIS_URL: "redis-h1ic17p47gr4df9i:6379" // TODO from resource
	}

	services: {
		web: {
			port: 5000
			kind: "http"
		}
	}

	// TODO add sync mount

	// Through adding a resource here, Namespace will
	// 1) instatiate a cache using Redis
	// 2) inject the configuration of the cache (e.g. endpoint) into the resource config of our Python server
	resources: {
		cache: {
			class:    "namespacelabs.dev/foundation/library/storage/redis:Database"
			provider: "namespacelabs.dev/foundation/library/oss/redis"
		}
	}
}
