# Go stack

This directory demonstrates how to model a Go application with Namespace.
For this example, we use MinIO to store S3 buckets.

`01-simple` is the simplest version, where each server is modeled separately and the stack is linked through `requires` statements.
Namespace injects a runtime configuration into each server. This allows a server to programatically consume the endpoint of its backends.

`02-withsecrets` models the MinIO server credentials as a Namespace secrets.
This allows the usage of an encrypted bundle to store the secret value.
Secret bundles can be created per server, or pinned for an entire workspace.
For this example, we use workspace-pinning - the password is stored in `workspace.secrets`.
You can reveal its value with `ns secrets reveal . --secret=namespacelabs.dev/examples/golang/02-withsecrets/minio:password` (we use an unencrypted bundle for demonstration purposes).

`03-withresources` consumes the S3 bucket as a Namespace resource provided by a shared MinIO provider.
The resource produces a typed instance object which provides credentials along with the endpoint for the Go server to consume.
Another advantage is that resources have their own lifetime modeling and initialization only happens once.
In the case of a bucket, this means that the bucket is created as part of the lifecycle and the Go server does not need to worry about it.
