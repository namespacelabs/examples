module: "namespacelabs.dev/examples"
requirements: {
	api: 35
}
dependency: {
	"namespacelabs.dev/foundation": {
		version: "b8992cacfcc7d7961b08d926a2e12963df193a34"
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
