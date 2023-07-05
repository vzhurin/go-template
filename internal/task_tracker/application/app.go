package application

import (
	"github.com/sirupsen/logrus"
	"github.com/vzhurin/template/internal/common/application"
	"github.com/vzhurin/template/internal/task_tracker/application/command"
	"github.com/vzhurin/template/internal/task_tracker/application/query"
	// TODO get rid of rtask
	rtask "github.com/vzhurin/template/internal/task_tracker/application/read_model/task"
	"github.com/vzhurin/template/internal/task_tracker/domain/task"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(
	taskRepository task.Repository,
	taskFinder rtask.Finder,
	logger logrus.FieldLogger,
) (*Application, func()) {
	return &Application{
			Commands: Commands{
				CreateTask: application.NewCommandLoggingDecorator[command.CreateTask](command.NewCreateTaskHandler(taskRepository), logger),
			},
			Queries: Queries{
				TaskByID: application.NewQueryLoggingDecorator[query.TaskByID, *rtask.Task](query.NewTaskByIDHandler(taskFinder), logger),
			},
		},
		func() { /* TODO add cleanup */ }
}

type Commands struct {
	CreateTask application.CommandHandler[command.CreateTask]
}

type Queries struct {
	TaskByID application.QueryHandler[query.TaskByID, *rtask.Task]
}
