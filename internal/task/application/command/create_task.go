package command

import "context"

type CreateTask struct{}

type CreateTaskHandler struct{}

func (h *CreateTaskHandler) Handle(ctx context.Context, cmd CreateTask) error {
	return nil
}

// TODO add decorators
