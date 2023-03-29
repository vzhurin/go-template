package query

import (
	"context"
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/common/application"
)

type GetTask struct {
	ID uuid.UUID
}

type getTaskHandler struct {
}

func NewGetTaskHandler() application.QueryHandler[GetTask, Task] {
	// TODO read model

	return &getTaskHandler{}
}

func (h *getTaskHandler) Handle(ctx context.Context, query GetTask) (Task, error) {
	return Task{
		ID:          uuid.Nil,
		Title:       "title",
		Description: "description",
		Assignee:    uuid.Nil,
		Status:      Completed,
		Estimation:  7,
	}, nil
}
