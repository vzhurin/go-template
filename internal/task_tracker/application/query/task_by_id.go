package query

import (
	"context"
	"github.com/vzhurin/template/internal/task_tracker/application/read_model/task"

	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/common/application"
)

type TaskByID struct {
	ID uuid.UUID
}

type taskByIDHandler struct {
	finder task.Finder
}

func NewTaskByIDHandler(finder task.Finder) application.QueryHandler[TaskByID, *task.Task] {
	return &taskByIDHandler{
		finder: finder,
	}
}

func (h *taskByIDHandler) Handle(ctx context.Context, query TaskByID) (*task.Task, error) {
	return h.finder.GetTaskByID(ctx, query.ID)
}
