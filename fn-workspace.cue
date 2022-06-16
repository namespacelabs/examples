module: "namespacelabs.dev/examples"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "6515a8626f5bd47f37db1c0d81dd198ac7443915"
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
