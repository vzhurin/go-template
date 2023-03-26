package command

import "context"

type Some struct{}

type SomeHandler struct{}

func (h *SomeHandler) Handle(ctx context.Context, cmd Some) error {
	return nil
}

// TODO add decorators
