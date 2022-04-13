// This file was automatically generated.
package main

import (
	"context"

	"namespacelabs.dev/examples/todos/api/todos"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/grpc/interceptors"
	"namespacelabs.dev/foundation/std/go/grpc/metrics"
	"namespacelabs.dev/foundation/std/go/grpc/server"
	"namespacelabs.dev/foundation/std/grpc"
	"namespacelabs.dev/foundation/std/grpc/deadlines"
	"namespacelabs.dev/foundation/std/grpc/logging"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
	"namespacelabs.dev/foundation/std/secrets"
	"namespacelabs.dev/foundation/universe/db/postgres/incluster"
	"namespacelabs.dev/foundation/universe/db/postgres/incluster/creds"
)

type ServerDeps struct {
	todos  todos.ServiceDeps
	trends trends.ServiceDeps
}

// This code uses type assertions for now. When go 1.18 is more widely deployed, it will switch to generics.
func PrepareDeps(ctx context.Context) (server *ServerDeps, err error) {
	di := core.MakeInitializer()

	di.Add(core.Provider{
		Package: "namespacelabs.dev/foundation/std/go/grpc/metrics",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps metrics.ExtensionDeps
			var err error

			if deps.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, nil); err != nil {
				return nil, err
			}

			return deps, err
		},
	})

	di.Add(core.Provider{
		Package: "namespacelabs.dev/foundation/std/monitoring/tracing",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps tracing.ExtensionDeps
			var err error

			if deps.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, nil); err != nil {
				return nil, err
			}

			return deps, err
		},
	})

	di.Add(core.Provider{
		Package: "namespacelabs.dev/foundation/universe/db/postgres/incluster/creds",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps creds.ExtensionDeps
			var err error
			// name: "postgres-password-file"
			p := &secrets.Secret{}
			core.MustUnwrapProto("ChZwb3N0Z3Jlcy1wYXNzd29yZC1maWxl", p)

			if deps.Password, err = secrets.ProvideSecret(ctx, p); err != nil {
				return nil, err
			}

			return deps, err
		},
	})

	di.Add(core.Provider{
		Package: "namespacelabs.dev/foundation/universe/db/postgres/incluster",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps incluster.ExtensionDeps
			var err error

			err = di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/universe/db/postgres/incluster/creds"},
				func(ctx context.Context, v interface{}) (err error) {

					if deps.Creds, err = creds.ProvideCreds(ctx, nil, v.(creds.ExtensionDeps)); err != nil {
						return err
					}
					return nil
				})
			if err != nil {
				return nil, err
			}

			{

				if deps.ReadinessCheck, err = core.ProvideReadinessCheck(ctx, nil); err != nil {
					return nil, err
				}
			}

			return deps, err
		},
	})

	di.Add(core.Provider{
		Package: "namespacelabs.dev/foundation/std/grpc/deadlines",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps deadlines.ExtensionDeps
			var err error

			if deps.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, nil); err != nil {
				return nil, err
			}

			return deps, err
		},
	})

	di.Add(core.Provider{
		Package: "namespacelabs.dev/examples/todos/api/todos",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps todos.ServiceDeps
			var err error

			err = di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/universe/db/postgres/incluster"},
				func(ctx context.Context, v interface{}) (err error) {
					// name: "todos"
					p := &incluster.Database{}
					core.MustUnwrapProto("CgV0b2Rvcw==", p)

					if deps.Db, err = incluster.ProvideDatabase(ctx, p, v.(incluster.ExtensionDeps)); err != nil {
						return err
					}
					return nil
				})
			if err != nil {
				return nil, err
			}

			err = di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/std/grpc/deadlines"},
				func(ctx context.Context, v interface{}) (err error) {
					// configuration: {
					//   service_name: "api.todos.TodosService"
					//   method_name: "List"
					//   maximum_deadline: 2
					// }
					// configuration: {
					//   service_name: "api.todos.TodosService"
					//   method_name: "GetRelatedData"
					//   maximum_deadline: 2
					// }
					p := &deadlines.Deadline{}
					core.MustUnwrapProto("CiMKFmFwaS50b2Rvcy5Ub2Rvc1NlcnZpY2USBExpc3QdAAAAQAotChZhcGkudG9kb3MuVG9kb3NTZXJ2aWNlEg5HZXRSZWxhdGVkRGF0YR0AAABA", p)

					if deps.Dl, err = deadlines.ProvideDeadlines(ctx, p, v.(deadlines.ExtensionDeps)); err != nil {
						return err
					}
					return nil
				})
			if err != nil {
				return nil, err
			}

			{
				// package_name: "namespacelabs.dev/examples/todos/api/trends"
				p := &grpc.Backend{}
				core.MustUnwrapProto("CituYW1lc3BhY2VsYWJzLmRldi9leGFtcGxlcy90b2Rvcy9hcGkvdHJlbmRz", p)

				if deps.TrendsConn, err = grpc.ProvideConn(ctx, p); err != nil {
					return nil, err
				}

				deps.Trends = trends.NewTrendsServiceClient(deps.TrendsConn)

			}

			return deps, err
		},
	})

	di.Add(core.Provider{
		Package: "namespacelabs.dev/foundation/std/grpc/logging",
		Do: func(ctx context.Context) (interface{}, error) {
			var deps logging.ExtensionDeps
			var err error

			if deps.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, nil); err != nil {
				return nil, err
			}

			return deps, err
		},
	})

	server = &ServerDeps{}
	di.AddInitializer(core.Initializer{
		PackageName: "",
		Do: func(ctx context.Context) error {

			err = di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/examples/todos/api/todos"},
				func(ctx context.Context, v interface{}) (err error) {
					server.todos = v.(todos.ServiceDeps)
					return nil
				})
			if err != nil {
				return err
			}

			err = di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/examples/todos/api/trends"},
				func(ctx context.Context, v interface{}) (err error) {
					server.trends = v.(trends.ServiceDeps)
					return nil
				})
			if err != nil {
				return err
			}

			return nil
		},
	})
	di.AddInitializer(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/metrics",
		Do: func(ctx context.Context) error {
			return di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/std/go/grpc/metrics"},
				func(ctx context.Context, v interface{}) (err error) {
					return metrics.Prepare(ctx, v.(metrics.ExtensionDeps))
				})
		},
	})

	di.AddInitializer(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/monitoring/tracing",
		Do: func(ctx context.Context) error {
			return di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/std/monitoring/tracing"},
				func(ctx context.Context, v interface{}) (err error) {
					return tracing.Prepare(ctx, v.(tracing.ExtensionDeps))
				})
		},
	})

	di.AddInitializer(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/grpc/deadlines",
		Do: func(ctx context.Context) error {
			return di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/std/grpc/deadlines"},
				func(ctx context.Context, v interface{}) (err error) {
					return deadlines.Prepare(ctx, v.(deadlines.ExtensionDeps))
				})
		},
	})

	di.AddInitializer(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/grpc/logging",
		Do: func(ctx context.Context) error {
			return di.Instantiate(ctx, core.Reference{Package: "namespacelabs.dev/foundation/std/grpc/logging"},
				func(ctx context.Context, v interface{}) (err error) {
					return logging.Prepare(ctx, v.(logging.ExtensionDeps))
				})
		},
	})

	return server, di.Init(ctx)
}

func WireServices(ctx context.Context, srv *server.Grpc, server *ServerDeps) {
	todos.WireService(ctx, srv, server.todos)
	srv.RegisterGrpcGateway(todos.RegisterTodosServiceHandler)
	trends.WireService(ctx, srv, server.trends)
}
