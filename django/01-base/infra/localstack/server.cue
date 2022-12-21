
import "namespacelabs.dev/foundation/library/oss/localstack/templates"

server: templates.#Server & {
	spec: {
		image: "localstack/localstack:latest"
		ingress: true
		dataVolume: {
			id: "localstack-server-django-todo"
			size: "1GiB"
		}
	}
}
