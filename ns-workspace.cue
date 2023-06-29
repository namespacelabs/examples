module: "namespacelabs.dev/examples"
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
		version: "2cfab843a3aa191483348578e6c7cd5c797a5974"
	}
}
