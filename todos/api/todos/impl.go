package todos

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/grpc/server"
	"namespacelabs.dev/go-ids"
)

type Service struct {
	DB     *pgxpool.Pool
	Trends trends.TrendsServiceClient
}

const timeout = 2 * time.Second

func addTodo(ctx context.Context, req *AddRequest, db *pgxpool.Pool) error {
	id := ids.NewSortableID()

	if _, err := db.Exec(ctx, "INSERT INTO todos_table (ID, Name) VALUES ($1, $2);", id, req.Name); err != nil {
		return fmt.Errorf("failed to add todo: %w", err)
	}

	return nil
}

func (svc *Service) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	if err := addTodo(ctx, req, svc.DB); err != nil {
		return nil, err
	}

	res := &AddResponse{}
	return res, nil
}

func removeTodo(ctx context.Context, req *RemoveRequest, db *pgxpool.Pool) error {
	// "Development" User Journey:
	// Uncomment next 3 lines.

	// if _, err := db.Exec(ctx, "DELETE FROM todos_table WHERE ID = $1;", req.Id); err != nil {
	// 	return fmt.Errorf("failed to remove todo: %w", err)
	// }

	return nil
}

func (svc *Service) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	if err := removeTodo(ctx, req, svc.DB); err != nil {
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
		return nil, fmt.Errorf("failed list todos: %w", err)
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
	todos, err := listTodos(svc.DB)
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
	name, err := getName(ctx, req.Id, svc.DB)
	if err != nil {
		return nil, err
	}

	pop, err := getPopularity(ctx, name, svc.Trends)
	if err != nil {
		return nil, err
	}

	return &GetRelatedDataResponse{Popularity: pop}, nil
}

func WireService(ctx context.Context, srv *server.Grpc, deps ServiceDeps) {
	RegisterTodosServiceServer(srv, &Service{DB: deps.Db, Trends: deps.Trends})
}
