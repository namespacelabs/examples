module: "namespacelabs.dev/examples"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "f014472068d807bd8780bd4d44f4fa538c792349"
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
