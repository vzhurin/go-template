package application

import (
	"github.com/sirupsen/logrus"
	"github.com/vzhurin/template/internal/common/application"
	"github.com/vzhurin/template/internal/task_tracker/application/command"
	"github.com/vzhurin/template/internal/task_tracker/application/query"
	"github.com/vzhurin/template/internal/task_tracker/domain/model/task"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(taskRepository task.Repository, logger logrus.FieldLogger) (*Application, func()) {
	return &Application{
			Commands: Commands{
				CreateTask: application.NewCommandLoggingDecorator[command.CreateTask](command.NewCreateTaskHandler(taskRepository), logger),
			},
			Queries: Queries{
				GetTask: application.NewQueryLoggingDecorator[query.GetTask, query.Task](query.NewGetTaskHandler(), logger),
			},
		},
		func() { /* TODO add cleanup */ }
}

type Commands struct {
	CreateTask application.CommandHandler[command.CreateTask]
}

type Queries struct {
	GetTask application.QueryHandler[query.GetTask, query.Task]
}
