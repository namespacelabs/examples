// This is a Namespace definition file.
// You can find a full syntax reference at https://namespace.so/docs/syntax-reference?utm_source=examples 
server: {
	name: "frontend"

	integration: web: {
		devPort: 5173

		// When adding a reference to the API server to the `backends` block, Namespace will
		// 1) add the API server to the deployed stack
		// 2) inject the configuration of API server (e.g. endpoint) into a src/config/backends.ns.js file that is accessible to the browser
		backends: {
			"api": "namespacelabs.dev/examples/multitier/01-simple/apibackend:webapi"
		}
	}
}
