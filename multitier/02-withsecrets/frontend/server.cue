server: {
	name: "frontend"

	integration: web: {
		service: "myweb"
		backends: {
			"api": "namespacelabs.dev/examples/multitier/02-withsecrets/apibackend:webapi"
		}
	}

	services: myweb: {
		// Default Vite port
		port: 5173
		kind: "http"

		ingress: internetFacing: true
	}
}
