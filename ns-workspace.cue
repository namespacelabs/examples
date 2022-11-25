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
		version: "a8ef9fe13e23568a9cb65bb2d793a9f8ac444c25"
	}
}
