# Multi-tier stack

This directory demonstrates how to model a multi-tier stack with Namespace.
For this example, we use a Vite frontend, an API server written in Go, and Postgres for the persistence layer.

`01-simple` is the simplest version, where each server is modeled separately and the stack is linked through `requires` statements.
Namespace injects a runtime configuration into each server. This allows a server to programatically consume the endpoint of its backends.

`02-withsecrets` models the Postgres server password as a Namespace secret.
For this example, we use a generated secret, since we don't care about the actual content, but only want to ensure that Postgres and our API backend use the same.

`03-withresources` consumes the database as a Namespace resource provided by a shared Postgres provider.
The resource produces a typed instance object which provides a password along with the endpoint for the Go backend to consume.
Another advantage is that resources have their own lifetime modeling and initialization only happens once.
In the case of a database, this means that the schema will be applied as part of the lifecycle and the Go backend does not need to worry about it.
