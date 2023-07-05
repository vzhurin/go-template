package task

import (
	"context"
	"github.com/google/uuid"
)

// TODO maybe Reader?
type Finder interface {
	// TODO maybe GetByID?
	GetTaskByID(ctx context.Context, id uuid.UUID) (*Task, error)
}
