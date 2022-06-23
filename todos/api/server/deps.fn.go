// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `ns generate`.

package main

import (
	"context"
	"namespacelabs.dev/examples/todos/api/todos"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/grpc/metrics"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/std/grpc/deadlines"
	"namespacelabs.dev/foundation/std/grpc/logging"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
	"namespacelabs.dev/foundation/std/monitoring/tracing/jaeger"
)

func RegisterInitializers(di *core.DependencyGraph) {
	di.AddInitializers(metrics.Initializers__so2f3v...)
	di.AddInitializers(tracing.Initializers__70o2mm...)
	di.AddInitializers(deadlines.Initializers__vbko45...)
	di.AddInitializers(logging.Initializers__16bc0q...)
	di.AddInitializers(jaeger.Initializers__33brri...)
}

func WireServices(ctx context.Context, srv server.Server, depgraph core.Dependencies) []error {
	var errs []error

	if err := depgraph.Instantiate(ctx, todos.Provider__i7grcp, func(ctx context.Context, v interface{}) error {
		todos.WireService(ctx, srv.Scope(todos.Package__i7grcp), v.(todos.ServiceDeps))
		return nil
	}); err != nil {
		errs = append(errs, err)
	}

	if err := depgraph.Instantiate(ctx, trends.Provider__021jd8, func(ctx context.Context, v interface{}) error {
		trends.WireService(ctx, srv.Scope(trends.Package__021jd8), v.(trends.ServiceDeps))
		return nil
	}); err != nil {
		errs = append(errs, err)
	}

	return errs
}
