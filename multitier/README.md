# Multi-tier stack

This directory demonstrates how to model a multi-tier stack with Namespace.
For this example, we use a Vite frontend, an API server written in Go, and Postgres for the persistence layer.

`01-simple` is the simplest version, where each server is modeled separately and the stack is linked through `requires` statements.
Namespace injects a runtime configuration into each server. This allows a server to programatically consume the endpoint of its backends.

`02-withresources` models the database as a Namespace resource.
This has the advantage that the resource can produce a typed instance object (e.g. provide a password along with the endpoint) for the Go backend to consume.
Another advantage is that resources have their own lifetime modeling and initialization only happens once.
In the case of a database, this means that the schema will be applied as part of the lifecycle and the Go backend does not need to worry about it.

`03-shared` is like `02-withresources` but uses a shared database definition from a separate directory (which is the typical usecase).
