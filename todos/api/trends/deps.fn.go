// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `fn generate`.

package trends

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/server"
)

// Dependencies that are instantiated once for the lifetime of the service.
type ServiceDeps struct {
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, server.Registrar, ServiceDeps)

var _ checkWireService = WireService

var (
	Package__021jd8 = &core.Package{
		PackageName: "namespacelabs.dev/examples/todos/api/trends",
	}

	Provider__021jd8 = core.Provider{
		Package:     Package__021jd8,
		Instantiate: makeDeps__021jd8,
	}
)

func makeDeps__021jd8(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ServiceDeps

	return deps, nil
}
