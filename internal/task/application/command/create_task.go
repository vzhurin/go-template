package command

import "context"

type CreateTask struct {
	// TODO
}

type CreateTaskHandler struct {
}

func NewCreateTaskHandler() *CreateTaskHandler {
	// TODO repository

	return &CreateTaskHandler{}
}

func (h *CreateTaskHandler) Handle(ctx context.Context, cmd CreateTask) error {
	// TODO
	return nil
}

// TODO add decorators
