package query

import (
	"context"
	"github.com/google/uuid"
)

type GetTask struct {
	ID uuid.UUID
}

type GetTaskHandler struct{}

func (h *GetTaskHandler) Handle(ctx context.Context, query GetTask) (Task, error) {
	// TODO repository

	return Task{
		ID:          uuid.Nil,
		Title:       "title",
		Description: "description",
		Assignee:    uuid.Nil,
		Status:      Completed,
		Estimation:  7,
	}, nil
}

// TODO add decorators
