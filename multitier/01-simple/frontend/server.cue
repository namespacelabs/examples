server: {
	name: "frontend"

	integration: web: service: "myweb"

	services: myweb: {
		// Default Vite port
		port: 5173
		kind: "http"

		ingress: internetFacing: true
	}
}
