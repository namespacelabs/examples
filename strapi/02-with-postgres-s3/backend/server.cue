server: {
	name: "strapi-backend"

	integration: "nodejs"

	env: {
		DATABASE_NAME: "bank"
		// Using a hard-coded passwords to simplify this example.
		// See multitier/02-withsecrets/postgres example for how to use a generated secret as the password.
		DATABASE_PASSWORD:    "DemoPasswordValue"
		S3_BUCKET:            "strapi-media"
		S3_ACCESS_KEY_ID:     "TestOnlyUser"
		S3_SECRET_ACCESS_KEY: "TestOnlyPassword"
	}

	services: backendapi: {
		port: 1337
		kind: "http"

		// The data API endpoint needs to be publicly available for client-side rendering
		ingress: internetFacing: true
	}

	requires: [
		"namespacelabs.dev/examples/strapi/02-with-postgres-s3/postgres",
		"namespacelabs.dev/examples/strapi/02-with-postgres-s3/minio",
	]
}
