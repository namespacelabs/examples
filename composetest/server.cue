server: {
	name: "composetest"

	integration: "dockerfile"

	env: {
		FLASK_DEBUG: "True"
	}

	services: {
		web: {
			port: 5000
			kind: "http"
		}
	}

	mounts: {
		"/code": configurable: {
			fromDir: "."
		}
	}

	requires: [
		"namespacelabs.dev/foundation/universe/db/redis/server",
	]
}
