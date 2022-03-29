import "namespacelabs.dev/foundation/std/fn"

server: fn.#Server & {
	id:        "dfbqvbafoqevrm6lm8o0"
	name:      "web-server"
	framework: "WEB"

	urlmap: [
		{path: "/", import: "namespacelabs.dev/examples/todos/web/ui"},
	]
}
