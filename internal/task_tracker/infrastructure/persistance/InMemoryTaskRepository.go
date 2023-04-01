package persistance

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/application/query"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
)

type InMemoryTaskRepository struct {
	storage map[string]map[string]interface{}
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		storage: make(map[string]map[string]interface{}),
	}
}

func (r *InMemoryTaskRepository) Get(ctx context.Context, id task.ID) (*task.Task, error) {
	t, ok := r.storage[id.String()]
	if !ok {
		return nil, fmt.Errorf("task with id %s not found", id)
	}

	return task.UnmarshalFromDatabase(
		t["id"].(string),
		t["title"].(string),
		t["description"].(string),
		t["assignee"].(string),
		t["status"].(string),
		t["estimation"].(uint64),
	), nil
}

func (r *InMemoryTaskRepository) Save(ctx context.Context, task *task.Task) error {
	r.storage[task.ID().String()] = task.MarshalToDatabase()

	fmt.Println(task.ID().String())

	return nil
}

func (r *InMemoryTaskRepository) TaskByID(ctx context.Context, id uuid.UUID) (*query.Task, error) {
	t, ok := r.storage[id.String()]
	if !ok {
		return nil, fmt.Errorf("task with id %s not found", id)
	}

	i, _ := uuid.Parse(t["id"].(string))
	a, _ := uuid.Parse(t["assignee"].(string))

	return &query.Task{
		ID:          i,
		Title:       t["title"].(string),
		Description: t["description"].(string),
		Assignee:    a,
		Status:      query.TaskStatus(t["status"].(string)),
		Estimation:  t["estimation"].(uint64),
	}, nil
}
