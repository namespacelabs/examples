import "namespacelabs.dev/foundation/library/oss/postgres/templates"

server: templates.#Server & {
	spec: {
		image: "postgres:13.9"
		dataVolume: {
			id: "postgres-server-data-django-todo"
			size: "1GiB"
		}
	}
}

secrets: {
	"password": {
		description: "Postgres server password"
		generate: {
			uniqueId:        "postgres-password"
			randomByteCount: 32
			format:          "FORMAT_BASE32"
		}
	}
}
