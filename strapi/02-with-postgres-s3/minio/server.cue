server: {
	name: "strapi-minio-server"

	image: "minio/minio@sha256:de46799fc1ced82b784554ba4602b677a71966148b77f5028132fc50adf37b1f"

	// Minio acts as an object store which requires a stateful deployment (more conservative update strategy). 
	class: "stateful"

	env: {
		// Disable update checking as self-update will never be used.
		MINIO_UPDATE:    "off"
		MINIO_ROOT_USER: "TestOnlyUser"
		// Using a hard-coded password to simplify this example.
		// See multitier/02-withsecrets/postgres example for how to use a generated secret as the password.
		MINIO_ROOT_PASSWORD: "TestOnlyPassword"
	}

	args: [
		"server",
		"/minio",
		"--address=:9000",
		"--console-address=:9001",
	]

	services: {
		minioapi: {
			port: 9000
			kind: "http"

			ingress: internetFacing: true
		}
		console: {
			port: 9001
			kind: "http"
		}
	}

	mounts: {
		"/minio": persistent: {
			// Unique volume identifier
			id:   "strapi-minio-server-data"
			size: "1GiB"
		}
	}
}
