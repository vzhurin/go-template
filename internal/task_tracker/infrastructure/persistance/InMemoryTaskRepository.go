package persistance

import (
	"context"
	"fmt"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
)

type InMemoryTaskRepository struct {
	storage map[task.ID]*task.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		storage: make(map[task.ID]*task.Task),
	}
}

func (r *InMemoryTaskRepository) Get(ctx context.Context, id task.ID) (*task.Task, error) {
	t, ok := r.storage[id]
	if !ok {
		return nil, fmt.Errorf("task with id %s not found", id)
	}

	return t, nil
}

func (r *InMemoryTaskRepository) Save(ctx context.Context, task *task.Task) error {
	r.storage[task.ID()] = task

	return nil
}
