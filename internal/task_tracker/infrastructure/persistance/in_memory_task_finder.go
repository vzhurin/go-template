package persistance

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task_tracker/application/read_model/task"
)

type InMemoryTaskFinder struct {
	storage map[string]map[string]interface{}
}

func NewInMemoryTaskFinder() *InMemoryTaskFinder {
	return &InMemoryTaskFinder{
		storage: make(map[string]map[string]interface{}),
	}
}

func (r *InMemoryTaskFinder) GetTaskByID(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	t, ok := r.storage[id.String()]
	if !ok {
		return nil, fmt.Errorf("task with id %s not found", id)
	}

	i, _ := uuid.Parse(t["id"].(string))
	a, _ := uuid.Parse(t["assignee"].(string))

	return &task.Task{
		ID:          i,
		Title:       t["title"].(string),
		Description: t["description"].(string),
		Assignee:    a,
		Status:      task.Status(t["status"].(string)),
		Estimation:  t["estimation"].(uint64),
	}, nil
}
