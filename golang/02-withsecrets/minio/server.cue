server: {
	name: "golang-secrets-minio-server"

	image: "minio/minio@sha256:de46799fc1ced82b784554ba4602b677a71966148b77f5028132fc50adf37b1f"

	// Minio acts as an object store which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// Disable update checking as self-update will never be used.
		MINIO_UPDATE: "off"

		// Instructs Namespace to inject the secrets as environment variables to the container.
		// See multitier/02-withsecrets/postgres/server.cue for an example injecting secrets into a mount.
		// These references point to secrets in the same package.
		// See golang/02-withsecrets/server/server.cue for an example using external secret references.
		MINIO_ROOT_USER: fromSecret:     ":user"
		MINIO_ROOT_PASSWORD: fromSecret: ":password"
	}

	args: [
		"server",
		"/minio",
		"--address=:9000",
		"--console-address=:9001",
	]

	services: {
		api: {
			port: 9000
			kind: "http"
		}
		console: {
			port: 9001
			kind: "http"
		}
	}

	mounts: {
		// Using an inline volume definition for brevity.
		// Mount points can also reference volumes by there package reference.
		// See multitier/02-withsecrets/postgres/server.cue for an example using a reference.
		"/minio": persistent: {
			// Unique volume identifier
			id:   "golang-secrets-minio-server-data"
			size: "10GiB"
		}
	}
}

secrets: {
	"user": {
		description: "Minio root user"
	}
	"password": {
		description: "Minio root password"
	}
}
