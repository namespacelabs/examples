import "namespacelabs.dev/foundation/std/fn"

test: fn.#Test & {
	name: "e2etest"

	binary: {
		from: go_package: "."
	}

	fixture: {
		sut: "namespacelabs.dev/examples/todos/api/server"
	}

	binary: {
		from: go_package: "."
	}
}
