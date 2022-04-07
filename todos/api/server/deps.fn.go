// This file was automatically generated.
package main

import (
	"context"

	"namespacelabs.dev/examples/todos/api/todos"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/grpc"
	"namespacelabs.dev/foundation/std/go/grpc/interceptors"
	"namespacelabs.dev/foundation/std/go/grpc/metrics"
	"namespacelabs.dev/foundation/std/go/grpc/server"
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

func PrepareDeps(ctx context.Context) (*ServerDeps, error) {
	var server ServerDeps
	var di core.DepInitializer
	var metrics0 metrics.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/interceptors",
		Instance:    "metrics0",
		Do: func(ctx context.Context) (err error) {
			if metrics0.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, "namespacelabs.dev/foundation/std/go/grpc/metrics", nil); err != nil {
				return err
			}
			return nil
		},
	})

	var tracing0 tracing.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/interceptors",
		Instance:    "tracing0",
		Do: func(ctx context.Context) (err error) {
			if tracing0.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, "namespacelabs.dev/foundation/std/monitoring/tracing", nil); err != nil {
				return err
			}
			return nil
		},
	})

	var creds0 creds.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/secrets",
		Instance:    "creds0",
		Do: func(ctx context.Context) (err error) {
			// name: "postgres-password-file"
			p := &secrets.Secret{}
			core.MustUnwrapProto("ChZwb3N0Z3Jlcy1wYXNzd29yZC1maWxl", p)

			if creds0.Password, err = secrets.ProvideSecret(ctx, "namespacelabs.dev/foundation/universe/db/postgres/incluster/creds", p); err != nil {
				return err
			}
			return nil
		},
	})

	var incluster0 incluster.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/universe/db/postgres/incluster/creds",
		Instance:    "incluster0",
		DependsOn:   []string{"creds0"}, Do: func(ctx context.Context) (err error) {
			if incluster0.Creds, err = creds.ProvideCreds(ctx, "namespacelabs.dev/foundation/universe/db/postgres/incluster", nil, creds0); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/core",
		Instance:    "incluster0",
		Do: func(ctx context.Context) (err error) {
			if incluster0.ReadinessCheck, err = core.ProvideReadinessCheck(ctx, "namespacelabs.dev/foundation/universe/db/postgres/incluster", nil); err != nil {
				return err
			}
			return nil
		},
	})

	var deadlines0 deadlines.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/interceptors",
		Instance:    "deadlines0",
		Do: func(ctx context.Context) (err error) {
			if deadlines0.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, "namespacelabs.dev/foundation/std/grpc/deadlines", nil); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/universe/db/postgres/incluster",
		Instance:    "server.todos",
		DependsOn:   []string{"incluster0"}, Do: func(ctx context.Context) (err error) {
			// name: "todos"
			p := &incluster.Database{}
			core.MustUnwrapProto("CgV0b2Rvcw==", p)

			if server.todos.Db, err = incluster.ProvideDatabase(ctx, "namespacelabs.dev/examples/todos/api/todos", p, incluster0); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/grpc/deadlines",
		Instance:    "server.todos",
		DependsOn:   []string{"deadlines0"}, Do: func(ctx context.Context) (err error) {
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

			if server.todos.Dl, err = deadlines.ProvideDeadlines(ctx, "namespacelabs.dev/examples/todos/api/todos", p, deadlines0); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc",
		Instance:    "server.todos",
		Do: func(ctx context.Context) (err error) {
			// package_name: "namespacelabs.dev/examples/todos/api/trends"
			// proto_typename: "TrendsService"
			p := &grpc.Conn{}
			core.MustUnwrapProto("CituYW1lc3BhY2VsYWJzLmRldi9leGFtcGxlcy90b2Rvcy9hcGkvdHJlbmRzEg1UcmVuZHNTZXJ2aWNl", p)

			if server.todos.TrendsConn, err = grpc.ProvideConn(ctx, "namespacelabs.dev/examples/todos/api/todos", p); err != nil {
				return err
			}

			server.todos.Trends = trends.NewTrendsServiceClient(server.todos.TrendsConn)
			return nil
		},
	})

	var logging0 logging.ExtensionDeps

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/interceptors",
		Instance:    "logging0",
		Do: func(ctx context.Context) (err error) {
			if logging0.Interceptors, err = interceptors.ProvideInterceptorRegistration(ctx, "namespacelabs.dev/foundation/std/grpc/logging", nil); err != nil {
				return err
			}
			return nil
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/go/grpc/metrics",
		DependsOn:   []string{"metrics0"},
		Do: func(ctx context.Context) error {
			return metrics.Prepare(ctx, metrics0)
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/monitoring/tracing",
		DependsOn:   []string{"tracing0"},
		Do: func(ctx context.Context) error {
			return tracing.Prepare(ctx, tracing0)
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/grpc/deadlines",
		DependsOn:   []string{"deadlines0"},
		Do: func(ctx context.Context) error {
			return deadlines.Prepare(ctx, deadlines0)
		},
	})

	di.Register(core.Initializer{
		PackageName: "namespacelabs.dev/foundation/std/grpc/logging",
		DependsOn:   []string{"logging0"},
		Do: func(ctx context.Context) error {
			return logging.Prepare(ctx, logging0)
		},
	})

	return &server, di.Wait(ctx)
}

func WireServices(ctx context.Context, srv *server.Grpc, server *ServerDeps) {
	todos.WireService(ctx, srv, server.todos)
	srv.RegisterGrpcGateway(todos.RegisterTodosServiceHandler)
	trends.WireService(ctx, srv, server.trends)
}
