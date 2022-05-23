import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/std/grpc"
	postgres "namespacelabs.dev/foundation/universe/db/postgres/incluster"
	"namespacelabs.dev/foundation/std/grpc/deadlines"
)

$proto: inputs.#Proto & {
	source: "service.proto"
}

service: fn.#Service & {
	framework: "GO_GRPC"

	instantiate: {
		trends: grpc.#Exports.Backend & {
			packageName: "namespacelabs.dev/examples/todos/api/trends"
		}

		db: postgres.#Exports.Database & {
			name:       "todos"
			schemaFile: inputs.#FromFile & {
				path: "schema.sql"
			}
		}

		dl: deadlines.#Exports.Deadlines & {
			configuration: [
				{serviceName: "api.todos.TodosService", methodName: "List", maximumDeadline:           2.0},
				{serviceName: "api.todos.TodosService", methodName: "GetRelatedData", maximumDeadline: 2.0},
			]
		}
	}

	ingress: "INTERNET_FACING"

	exportService:        $proto.services.TodosService
	exportServicesAsHttp: true
}
