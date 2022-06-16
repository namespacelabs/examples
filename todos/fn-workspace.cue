module: "namespacelabs.dev/examples/todos"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "ce1b880b5907de5db299625b9ace1822b29e11e6"
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
