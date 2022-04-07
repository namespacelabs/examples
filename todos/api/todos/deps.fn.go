// This file was automatically generated.
package todos

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/grpc/server"
	"namespacelabs.dev/foundation/std/grpc/deadlines"
)

type ServiceDeps struct {
	Db         *pgxpool.Pool
	Dl         *deadlines.DeadlineRegistration
	Trends     trends.TrendsServiceClient
	TrendsConn *grpc.ClientConn
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, *server.Grpc, ServiceDeps)

var _ checkWireService = WireService
