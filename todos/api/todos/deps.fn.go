// This file was automatically generated.
package todos

import (
	"context"

	_ "namespacelabs.dev/foundation/std/go/grpc/metrics"
	_ "namespacelabs.dev/foundation/std/monitoring/tracing"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/grpc/server"
)

type ServiceDeps struct {
	Db         *pgxpool.Pool
	Trends     trends.TrendsServiceClient
	TrendsConn *grpc.ClientConn
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, *server.Grpc, ServiceDeps)

var _ checkWireService = WireService
