module: "namespacelabs.dev/examples"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "39d109835ec70a838f8f4bfe4fd07a47c17b100a"
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
