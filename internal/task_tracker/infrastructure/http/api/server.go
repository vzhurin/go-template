package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	ht "github.com/vzhurin/template/internal/common/infrastructure/http"
	"github.com/vzhurin/template/internal/task_tracker/application"
	"github.com/vzhurin/template/internal/task_tracker/application/command"
	"github.com/vzhurin/template/internal/task_tracker/application/query"
)

type Server struct {
	app    *application.Application
	logger logrus.FieldLogger
}

func NewServer(app *application.Application, logger logrus.FieldLogger) *Server {
	return &Server{
		app:    app,
		logger: logger,
	}
}

func (s *Server) EstimateTask(w http.ResponseWriter, r *http.Request, taskID UUID) {

}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ht.InternalError(err.Error(), w, r)
	}

	task := &PostTask{}
	if err := json.Unmarshal(body, task); err != nil {
		ht.BadRequest(err.Error(), w, r)
	}

	cmd := command.CreateTask{
		Title:       task.Title,
		Description: task.Description,
	}

	err = s.app.Commands.CreateTask.Handle(r.Context(), cmd)
	if err != nil {
		ht.BadRequest(err.Error(), w, r)
	}

}

func (s *Server) GetTask(w http.ResponseWriter, r *http.Request, taskID UUID) {
	task, err := s.app.Queries.TaskByID.Handle(r.Context(), query.TaskByID{ID: taskID})
	if err != nil {
		s.logger.WithError(err).Warning("task not found")
		ht.NotFound(err.Error(), w, r)
		return
	}

	payload, err := json.Marshal(Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Assignee:    task.Assignee,
		Status:      TaskStatus(task.Status),
		Estimation:  task.Estimation,
	})

	if err != nil {
		s.logger.WithError(err).Error("marshaling failed")
		ht.InternalError(err.Error(), w, r)

		return
	}

	_, err = w.Write(payload)

	if err != nil {
		s.logger.WithError(err).Error("response failed")
		ht.InternalError(err.Error(), w, r)
		return
	}
}

func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request, taskID UUID) {

}

func (s *Server) AssignParticipantToTask(w http.ResponseWriter, r *http.Request, taskID UUID) {

}

func (s *Server) TransitTaskStatus(w http.ResponseWriter, r *http.Request, taskID UUID) {

}
