package task

import "context"

type Repository interface {
	Get(ctx context.Context, id ID) (*Task, error)
	Save(ctx context.Context, task *Task) error
}
