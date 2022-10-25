// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
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
		// This mount point references a volume in the same package.
		// For external volume references, use "example.com/pkg/path:volumeName".
		"/postgres/data": ":data"

		// Alternative syntax: inline volume definition.
		"/postgres/secrets": configurable: {
			contents: {
				// Instructs Namespace to inject the secrets as file in the mount.
				// See golang/02-withsecrets/minio/server.cue for an example injecting secrets as environment variables.
				"password": fromSecret: ":password"
			}
		}
	}
}

volumes: {
	"data": persistent: {
		// Unique volume identifier
		id:   "multitier-secrets-postgres-server-data"
		size: "10GiB"
	}
}

secrets: {
	"password": {
		description: "Postgres server password"
		generate: {
			uniqueId:        "multitier-secrets-postgres-password"
			randomByteCount: 32
			format:          "FORMAT_BASE32"
		}
	}
}
