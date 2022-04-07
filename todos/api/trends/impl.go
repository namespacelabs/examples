package trends

import (
	"context"
	"hash/fnv"

	"namespacelabs.dev/foundation/std/go/grpc/server"
)

type Service struct {
}

func getPopularity(name string) uint32 {
	// Produce some fake but stable popularity values.
	h := fnv.New32a()
	h.Write([]byte(name))
	pop := (h.Sum32() % 5) + 1

	return pop
}

func (svc *Service) GetTrends(ctx context.Context, req *GetTrendsRequest) (*GetTrendsResponse, error) {
	pop := getPopularity(req.Name)
	res := &GetTrendsResponse{Popularity: pop}
	return res, nil
}

func WireService(ctx context.Context, srv *server.Grpc, deps ServiceDeps) {
	RegisterTrendsServiceServer(srv, &Service{})
}
