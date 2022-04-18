package todos

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"namespacelabs.dev/examples/todos/api/trends"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/go-ids"
)

type Service struct {
	DB     *pgxpool.Pool
	Trends trends.TrendsServiceClient
}

func (svc *Service) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	if err := addTodo(ctx, req, svc.DB); err != nil {
		return nil, err
	}

	res := &AddResponse{}
	return res, nil
}

func (svc *Service) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	if err := removeTodo(ctx, req, svc.DB); err != nil {
		return nil, err
	}

	res := &RemoveResponse{}
	return res, nil
}

func (svc *Service) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	todos, err := fetchTodosList(ctx, svc.DB)
	if err != nil {
		return nil, err
	}

	res := &ListResponse{
		Items: todos,
	}
	return res, nil
}

func (svc *Service) GetRelatedData(ctx context.Context, req *GetRelatedDataRequest) (*GetRelatedDataResponse, error) {
	name, err := fetchName(ctx, req.Id, svc.DB)
	if err != nil {
		return nil, err
	}

	pop, err := computePopularity(ctx, name, svc.Trends)
	if err != nil {
		return nil, err
	}

	return &GetRelatedDataResponse{Popularity: pop}, nil
}

func addTodo(ctx context.Context, req *AddRequest, db *pgxpool.Pool) error {
	id := ids.NewSortableID()

	if _, err := db.Exec(ctx, "INSERT INTO todos_table (ID, Name) VALUES ($1, $2);", id, req.Name); err != nil {
		return fmt.Errorf("failed to add todo: %w", err)
	}

	return nil
}

func removeTodo(ctx context.Context, req *RemoveRequest, db *pgxpool.Pool) error {
	// "Development" User Journey:
	// Uncomment next 3 lines.

	// if _, err := db.Exec(ctx, "DELETE FROM todos_table WHERE ID = $1;", req.Id); err != nil {
	// 	return fmt.Errorf("failed to remove todo: %w", err)
	// }

	return nil
}

func fetchTodosList(ctx context.Context, db *pgxpool.Pool) ([]*TodoItem, error) {
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

func fetchName(ctx context.Context, id string, db *pgxpool.Pool) (string, error) {
	var name string
	if err := db.QueryRow(ctx, "SELECT Name FROM todos_table WHERE ID = $1;", id).Scan(&name); err != nil {
		log.Printf("failed to read todos: %v", err)
		return "", err
	}

	return name, nil
}

func computePopularity(ctx context.Context, name string, client trends.TrendsServiceClient) (uint32, error) {
	req := &trends.GetTrendsRequest{
		Name: name,
	}
	resp, err := client.GetTrends(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.Popularity, nil
}

func WireService(ctx context.Context, srv server.Registrar, deps ServiceDeps) {
	RegisterTodosServiceServer(srv, &Service{DB: deps.Db, Trends: deps.Trends})
}
