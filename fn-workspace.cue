module: "namespacelabs.dev/examples"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "0e1f6784e25d9d8b5454f7a10ff1a18329edee2b"
	}
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
