server: {
	name: "nextjsserver"

	integration: "nodejs"

	env: {
		"PORT": "4000"
	}

	services: {
		webapi: {
			port: 4000
			kind: "http"

			ingress: {
				internetFacing: true
				httpRoutes: "*": ["/"]
			}
		}
	}
}
