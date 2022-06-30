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
		version: "92a0405725f5af65ad1ce38809a7edca0874465e"
	}
}
