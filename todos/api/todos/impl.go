package todos

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/proto"
	"namespacelabs.dev/examples/todo-app/api/trends"
	"namespacelabs.dev/foundation/std/go/grpc/server"
	"namespacelabs.dev/go-ids"
)

type Service struct {
	deps ServiceDeps
}

const timeout = 2 * time.Second

func logRequest(req proto.Message) {
	log.Printf("new %s request: %+v\n", req.ProtoReflect().Descriptor().FullName(), req)
}

func addTodo(req *AddRequest, db *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	id := ids.NewSortableID()

	if _, err := db.Exec(ctx, "INSERT INTO todos_table (ID, Name) VALUES ($1, $2);", id, req.Name); err != nil {
		return fmt.Errorf("failed to add todo: %w", err)
	}

	return nil
}

func (svc *Service) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	logRequest(req)

	if err := addTodo(req, svc.deps.Db); err != nil {
		return nil, err
	}

	res := &AddResponse{}
	return res, nil
}

func removeTodo(req *RemoveRequest, db *pgxpool.Pool) error {
	// "Development" User Journey:
	// Uncomment next 5 lines.

	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()

	// if _, err := db.Exec(ctx, "DELETE FROM todos_table WHERE ID = $1;", req.Id); err != nil {
	// 	return fmt.Errorf("failed to remove todo: %w", err)
	// }

	return nil
}

func (svc *Service) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	logRequest(req)

	if err := removeTodo(req, svc.deps.Db); err != nil {
		return nil, err
	}

	res := &RemoveResponse{}
	return res, nil
}

func listTodos(db *pgxpool.Pool) ([]*TodoItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	rows, err := db.Query(ctx, "SELECT ID, Name FROM todos_table;")
	if err != nil {
		return nil, fmt.Errorf("failed to remove todo: %w", err)
	}
	defer rows.Close()

	var res []*TodoItem
	for rows.Next() {
		todo := &TodoItem{}
		err = rows.Scan(&todo.Id, &todo.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, todo)
	}

	return res, nil
}

func (svc *Service) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	logRequest(req)

	todos, err := listTodos(svc.deps.Db)
	if err != nil {
		return nil, err
	}

	res := &ListResponse{
		Items: todos,
	}
	return res, nil
}

func getName(ctx context.Context, id string, db *pgxpool.Pool) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var name string
	if err := db.QueryRow(ctx, "SELECT Name FROM todos_table WHERE ID = $1;", id).Scan(&name); err != nil {
		log.Printf("failed to read todos: %v", err)
		return "", err
	}

	return name, nil
}

func getPopularity(ctx context.Context, name string, client trends.TrendsServiceClient) (uint32, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req := &trends.GetTrendsRequest{
		Name: name,
	}
	resp, err := client.GetTrends(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.Popularity, nil

}

func (svc *Service) GetRelatedData(ctx context.Context, req *GetRelatedDataRequest) (*GetRelatedDataResponse, error) {
	logRequest(req)

	name, err := getName(ctx, req.Id, svc.deps.Db)
	if err != nil {
		return nil, err
	}

	pop, err := getPopularity(ctx, name, svc.deps.Trends)
	if err != nil {
		return nil, err
	}

	response := &GetRelatedDataResponse{Popularity: pop}

	log.Printf("will reply with: %+v\n", response)

	return response, nil
}

func WireService(ctx context.Context, srv *server.Grpc, deps ServiceDeps) {
	RegisterTodosServiceServer(srv, &Service{deps: deps})
}
