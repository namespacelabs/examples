import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/std/web/http"
)

$apiServer: inputs.#Server & {
	packageName: "namespacelabs.dev/examples/todos/api/server"
}

service: fn.#Service & {
	framework: "WEB"

	instantiate: {
		apiBackend: http.#Exports.Backend & {
			with: {
				endpointOwner: $apiServer.packageName
				serviceName: "grpc-gateway"
			}
		}
	}
}

extend: fn.#Extend & {
	stack: {
		append: [$apiServer]
	}
}