server: {
	name: "multitier-secrets-postgres-server"

	image: "postgres:14.0@sha256:db927beee892dd02fbe963559f29a7867708747934812a80f83bff406a0d54fd"

	// Postgres mounts a persistent volume which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// PGDATA may not be a mount point but only a subdirectory.
		PGDATA:                 "/postgres/data/pgdata"
		POSTGRES_DB:            "todos"
		POSTGRES_PASSWORD_FILE: "/postgres/secrets/password"
	}

	services: "postgres": {
		port: 5432
		kind: "tcp"
	}

	mounts: {
		"/postgres/data":    "data"
		"/postgres/secrets": "secrets"
	}
}

volumes: {
	"data": persistent: {
		// Unique volume identifier
		id:   "multitier-secrets-postgres-server-data"
		size: "10GiB"
	}
	"secrets": configurable: {
		contents: {
			"password": fromSecret: "password"
		}
	}
}

secrets: {
	"password": {
		description: "Postgres server password"
	}
}