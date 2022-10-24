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
		version: "f8ad005f1e92ceb88118d2296840703b0ddcbbe8"
	}
}
