// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
server: {
	name: "strapi-postgres-server"

	image: "postgres:14.0@sha256:db927beee892dd02fbe963559f29a7867708747934812a80f83bff406a0d54fd"

	// Postgres mounts a persistent volume which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// PGDATA may not be a mount point but only a subdirectory.
		PGDATA:      "/postgres/data/pgdata"
		POSTGRES_DB: "bank"
		// Using a hard-coded password to simplify this example.
		// See multitier/02-withsecrets/postgres example for how to use a generated secret as the password.
		POSTGRES_PASSWORD: "DemoPasswordValue"
	}

	services: "postgres": {
		port: 5432
		kind: "tcp"
	}

	mounts: {
		"/postgres/data": persistent: {
			// Unique volume identifier
			id:   "strapi-postgres-server-data"
			size: "10GiB"
		}
	}
}
