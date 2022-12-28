server: {
	name: "java-spring-boot-demo"

	integration: "dockerfile"

	services: {
		webapi: {
			port: 8080
			kind: "http"
		}
	}
}
