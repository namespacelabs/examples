# Next.js application

This directory demonstrates how to model a Next.js application with Namespace.
For this example, we use Postgres as a persistence layer.

`01-simple` is the simplest version, where each server is modeled separately and the Postgres is linked through a `requires` statement in the Next.js server definition.
Namespace injects a runtime configuration into the Next.js server. This allows it to programatically consume the endpoint of its Postgres backend.

`01-withsecrets` models the Postgres server password as a Namespace secret.
This allows the usage of an encrypted bundle to store the secret value.
Secret bundles can be created per server, or pinned for an entire workspace.
For this example, we use workspace-pinning - the password is stored in `workspace.secrets`.
You can reveal its value with `ns secrets reveal . --secret=namespacelabs.dev/examples/nextjs/02-withsecrets/postgres:password` (we use an unencrypted bundle for demonstration purposes).

`03-withresources` models the database as a Namespace resource.
This has the advantage that the resource can produce a typed instance object (e.g. provide a password along with the endpoint) for the Next.js server to consume.
Another advantage is that resources have their own lifetime modeling and initialization only happens once.
In the case of a database, this means that the schema will be applied as part of the lifecycle and the Next.js server does not need to worry about it.

`03-shared` is like `03-withresources` but uses a shared database definition from a separate directory (which is the typical usecase).
