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

func NewApplication(
	taskRepository task.Repository,
	taskByIDReadModel query.TaskByIDReadModel,
	logger logrus.FieldLogger,
) (*Application, func()) {
	return &Application{
			Commands: Commands{
				CreateTask: application.NewCommandLoggingDecorator[command.CreateTask](command.NewCreateTaskHandler(taskRepository), logger),
			},
			Queries: Queries{
				TaskByID: application.NewQueryLoggingDecorator[query.TaskByID, *query.Task](query.NewTaskByIDHandler(taskByIDReadModel), logger),
			},
		},
		func() { /* TODO add cleanup */ }
}

type Commands struct {
	CreateTask application.CommandHandler[command.CreateTask]
}

type Queries struct {
	TaskByID application.QueryHandler[query.TaskByID, *query.Task]
}
