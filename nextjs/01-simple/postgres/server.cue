server: {
	name: "nextjs-simple-postgres-server"

	image: "postgres:14.0@sha256:db927beee892dd02fbe963559f29a7867708747934812a80f83bff406a0d54fd"

	// Postgres mounts a persistent volume which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// PGDATA may not be a mount point but only a subdirectory.
		PGDATA:            "/postgres/data/pgdata"
		POSTGRES_DB:       "nextjs"
		POSTGRES_PASSWORD: "DemoPasswordValue"
	}

	services: "postgres": {
		port: 5432
		kind: "tcp"
	}

	mounts: {
		"/postgres/data": persistent: {
			// Unique volume identifier
			id:   "nextjs-simple-postgres-server-data"
			size: "10GiB"
		}
	}
}
