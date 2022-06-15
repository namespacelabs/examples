module: "namespacelabs.dev/examples/todos"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "e6e6beab42ac124163a40f7a33af290a7b5292f0"
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
