server: {
	name: "golang-withresources-minio-server"

	image: "minio/minio@sha256:de46799fc1ced82b784554ba4602b677a71966148b77f5028132fc50adf37b1f"

	// Minio acts as an object store which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// Disable update checking as self-update will never be used.
		MINIO_UPDATE: "off"
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
		"/minio": "data"
	}
}

volumes: {
	"data": persistent: {
		// Unique volume identifier
		id:   "golang-withresources-minio-server-data"
		size: "10GiB"
	}
}
