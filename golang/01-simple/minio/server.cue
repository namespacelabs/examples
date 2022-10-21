server: {
	name: "golang-simple-minio-server"

	image: "minio/minio@sha256:de46799fc1ced82b784554ba4602b677a71966148b77f5028132fc50adf37b1f"

	// Minio acts as an object store which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// Disable update checking as self-update will never be used.
		MINIO_UPDATE: "off"

		// Using hard-coded credentials to simplify this example.
		// See golang/02-withsecrets/ for an example using managed secrets.
		MINIO_ROOT_USER:     "TestOnlyUser"
		MINIO_ROOT_PASSWORD: "TestOnlyPassword"
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
		"/minio": persistent: {
			// Unique volume identifier
			id:   "golang-simple-minio-server-data"
			size: "10GiB"
		}
	}
}
