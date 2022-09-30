server: {
	name: "nextjsserver"

	integration: "nodejs"

	services: {
		webapi: {
			port: 3000
			kind: "http"

			ingress: {
				internetFacing: true
				httpRoutes: "*": ["/"]
			}
		}
	}
}