module: "namespacelabs.dev/examples/todos"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "ed1a2e2b377f2c1fe641097558352ac33f35b723"
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
