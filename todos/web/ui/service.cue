import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/std/web/http"
)

service: fn.#Service & {
	framework: "WEB"

	instantiate: {
		apiBackend: http.#Exports.Backend & {
			endpointOwner: "namespacelabs.dev/examples/todos/api/server"
			serviceName:   "grpc-gateway"
		}
	}
}
