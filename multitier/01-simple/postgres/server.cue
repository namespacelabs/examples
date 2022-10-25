// This is a Namespace definition file.
// You can find a full syntax reference at https://docs.namespace.so/reference?utm_source=examples 
server: {
	name: "multitier-simple-postgres-server"

	image: "postgres:14.0@sha256:db927beee892dd02fbe963559f29a7867708747934812a80f83bff406a0d54fd"

	// Postgres mounts a persistent volume which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// PGDATA may not be a mount point but only a subdirectory.
		PGDATA:      "/postgres/data/pgdata"
		POSTGRES_DB: "todos"

		// Using a hard-coded password to simplify this example.
		// See multitier/02-withsecrets/ for an example using generated secrets.
		POSTGRES_PASSWORD: "DemoPasswordValue"
	}

	services: "postgres": {
		port: 5432
		kind: "tcp"
	}

	mounts: {
		// This mount point references a volume in the same package.
		// For external volume references, use "example.com/pkg/path:volumeName".
		// See golang/01-simple/minio/server.cue for an example using an inline volume definition.
		"/postgres/data": ":data"
	}
}

volumes: {
	"data": persistent: {
		// Unique volume identifier
		id:   "multitier-simple-postgres-server-data"
		size: "10GiB"
	}
}
