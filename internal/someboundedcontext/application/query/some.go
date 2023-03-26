package query

import "context"

type Some struct{}

type Thing struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

type SomeHandler struct{}

func (h *SomeHandler) Handle(ctx context.Context, query Some) ([]Thing, error) {
	return []Thing{{Field1: "one", Field2: 2}}, nil
}

// TODO add decorators
