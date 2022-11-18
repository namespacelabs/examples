// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "nextjs-secrets-postgres-server"

	image: "postgres:14.0@sha256:db927beee892dd02fbe963559f29a7867708747934812a80f83bff406a0d54fd"

	// Postgres mounts a persistent volume which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// PGDATA may not be a mount point but only a subdirectory.
		PGDATA:                 "/postgres/data/pgdata"
		POSTGRES_DB:            "nextjs"
		POSTGRES_PASSWORD_FILE: "/postgres/secrets/password"
	}

	services: "postgres": {
		port: 5432
		kind: "tcp"
	}

	mounts: {
		// Using inline volume definitions for brevity.
		// Mount points can also reference volumes by their package reference.
		// See multitier/02-withsecrets/postgres/server.cue for an example using a reference.
		"/postgres/data": persistent: {
			// Unique volume identifier
			id:   "nextjs-secrets-postgres-server-data"
			size: "10GiB"
		}
		"/postgres/secrets": configurable: {
			contents: {
				// Instructs Namespace to inject the secrets as file in the mount.
				// See golang/02-withsecrets/minio/server.cue for an example injecting secrets as environment variables.
				"password": fromSecret: ":password"
			}
		}
	}
}

secrets: {
	"password": {
		description: "Postgres server password"
	}
}
