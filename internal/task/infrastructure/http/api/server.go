package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/vzhurin/template/internal/task/application"
	"github.com/vzhurin/template/internal/task/application/query"
	"io"
	"net/http"
)

type Server struct {
	app *application.Application
}

func NewServer(app *application.Application) *Server {
	return &Server{
		app: app,
	}
}

func (s *Server) EstimateTask(w http.ResponseWriter, r *http.Request, taskID UUID) {

}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// TODO
	}

	task := &PostTask{}
	if err := json.Unmarshal(body, task); err != nil {
		// TODO
	}

	fmt.Printf("%#v", task)
}

func (s *Server) GetTask(w http.ResponseWriter, r *http.Request, taskID UUID) {
	task, err := s.app.Queries.GetTask.Handle(r.Context(), query.GetTask{ID: uuid.Nil})
	if err != nil {
		// TODO
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
		// TODO
	}

	_, err = w.Write(payload)

	if err != nil {
		// TODO
	}
}

func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request, taskID UUID) {

}

func (s *Server) AssignParticipantToTask(w http.ResponseWriter, r *http.Request, taskID UUID) {

}

func (s *Server) TransitTaskStatus(w http.ResponseWriter, r *http.Request, taskID UUID) {

}
