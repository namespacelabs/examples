import "namespacelabs.dev/foundation/std/fn"

server: fn.#Server & {
	id:        "9r5mlstodp2kacg51e0g"
	name:      "api-backend"
	framework: "GO_GRPC"

	import: [
		"namespacelabs.dev/foundation/std/go/grpc/gateway",
		"namespacelabs.dev/examples/todo-app/api/todos",
		"namespacelabs.dev/examples/todo-app/api/trends",
	]
}
