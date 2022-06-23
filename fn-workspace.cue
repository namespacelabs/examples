module: "namespacelabs.dev/examples"
requirements: {
	api: 35
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
		version: "c08091abb53acfc4dfb03fa6f7c5aac42cb4a609"
	}
}
