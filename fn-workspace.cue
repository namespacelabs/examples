module: "namespacelabs.dev/examples"
requirements: {
	api: 41
}
environment: {
	dev: {
		runtime: "kubernetes"
		purpose: "DEVELOPMENT"
	}
	staging: {
		runtime: "kubernetes"
		purpose: "PRODUCTION"
	}
	prod: {
		runtime: "kubernetes"
		purpose: "PRODUCTION"
	}
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "54fdfb78b0e1efeb7015d30e19fcba65ad0c2915"
	}
}
