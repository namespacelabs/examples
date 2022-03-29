import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/std/grpc"
	"namespacelabs.dev/foundation/universe/db/postgres/incluster"
)

$proto: inputs.#Proto & {
	source: "service.proto"
}

$backend: grpc.#Backend & {
	packageName: "namespacelabs.dev/examples/todos/api/trends"
}

service: fn.#Service & {
	instantiate: {
		$backend.instances

		db: incluster.#Exports.Database & {
			with: {
				name:       "todos"
				schemaFile: inputs.#FromFile & {
					path: "schema.sql"
				}
			}
		}
	}

	ingress: "INTERNET_FACING"

	exportService:        $proto.services.TodosService
	exportServicesAsHttp: true
}
