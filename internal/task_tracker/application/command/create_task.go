package command

import (
	"context"
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/common/application"
	"github.com/vzhurin/template/internal/task_tracker/domain/task"
)

type CreateTask struct {
	Title       string
	Description string
}

type createTaskHandler struct {
	repository task.Repository
}

func NewCreateTaskHandler(repository task.Repository) application.CommandHandler[CreateTask] {
	return &createTaskHandler{
		repository: repository,
	}
}

func (h *createTaskHandler) Handle(ctx context.Context, cmd CreateTask) error {
	id, _ := task.NewID(uuid.NewString())

	title, err := task.NewTitle(cmd.Title)
	if err != nil {
		return err
	}

	description, err := task.NewDescription(cmd.Description)
	if err != nil {
		return err
	}

	t := task.NewTask(id, title, description)

	_, err = h.repository.Save(ctx, t)
	if err != nil {
		return err
	}

	return nil
}
