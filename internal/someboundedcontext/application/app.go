package application

import (
	"template/internal/someboundedcontext/application/command"
	"template/internal/someboundedcontext/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication() (*Application, func()) {
	return &Application{}, func() {}
}

type Commands struct {
	Some command.SomeHandler
}

type Queries struct {
	Some query.SomeHandler
}
