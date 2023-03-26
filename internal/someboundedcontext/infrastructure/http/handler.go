package http

import (
	"encoding/json"
	"net/http"
	"template/internal/someboundedcontext/application"
	"template/internal/someboundedcontext/application/command"
	"template/internal/someboundedcontext/application/query"
)

type Handler struct {
	app *application.Application
}

func NewHandler(app *application.Application) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) CreateSome(w http.ResponseWriter, r *http.Request) {
	err := h.app.Commands.Some.Handle(r.Context(), command.Some{})
	if err != nil {

	}
}

func (h *Handler) GetSome(w http.ResponseWriter, r *http.Request) {
	things, err := h.app.Queries.Some.Handle(r.Context(), query.Some{})
	if err != nil {

	}

	payload, err := json.Marshal(things)
	if err != nil {

	}

	w.Write(payload)
}
