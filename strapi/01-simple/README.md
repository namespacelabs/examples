This example demonstrates how to package an existing complex server stack: a Strapi backend + a
Next.js frontend.

The code skeleton is generated using `yarn create strapi-starter my-project next-blog`. The default
Strapi config is used: committed media and sqllite as the database.

To make it work with Namespace, the following changes were made:

- Added `.cue` configuration files.
- Added `dev` script to `strapi/backend`, this is what Namespace calls in dev mode.
- In the Next.js frontend, configured the URL of the Strapi backend to be read from the Namespace
  config (on the server) or from the generated `backends.ns.js` (in the browser).
- The Next.js image caching loader has been disabled, it is not fully supported by Namespace yet.
- Static site generation is not supported for prod builds: it requires the Strapi backend to be
  available to the Next.js frontend at build time, which is not possible: the builds should be
  hermetic. So the tests are configured to use the `dev` version, in `strapi/frontend/package.json`:
  - The `build` script has been removed.
  - The `start` script changed from `next start` to `next dev`.
