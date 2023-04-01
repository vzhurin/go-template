package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/common/application"
)

type TaskByID struct {
	ID uuid.UUID
}

type taskByIDHandler struct {
	readModel TaskByIDReadModel
}

func NewTaskByIDHandler(readModel TaskByIDReadModel) application.QueryHandler[TaskByID, *Task] {
	return &taskByIDHandler{
		readModel: readModel,
	}
}

func (h *taskByIDHandler) Handle(ctx context.Context, query TaskByID) (*Task, error) {
	return h.readModel.TaskByID(ctx, query.ID)
}

type TaskByIDReadModel interface {
	TaskByID(ctx context.Context, id uuid.UUID) (*Task, error)
}
