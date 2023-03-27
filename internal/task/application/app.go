package application

import (
	"github.com/vzhurin/template/internal/task/application/command"
	"github.com/vzhurin/template/internal/task/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication() (*Application, func()) {
	return &Application{}, func() { /* TODO add cleanup */ }
}

type Commands struct {
	CreateTask command.CreateTaskHandler
}

type Queries struct {
	GetTask query.GetTaskHandler
}
