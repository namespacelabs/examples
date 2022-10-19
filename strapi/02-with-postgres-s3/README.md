This example demonstrates how to package an existing complex server stack: a Strapi backend + a
Next.js frontend.

The code skeleton is generated using `yarn create strapi-starter my-project next-blog`.

Differences from the `01-simple` example:

- Using an in-cluster Postgres database insted of sqllite.
- Using Minio as S3 storage for media, instead of the local filesystem.
