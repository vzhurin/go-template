package http

import "github.com/go-chi/chi/v5"

func NewRouter(handler *Handler) chi.Router {
	router := chi.NewRouter()

	router.Get("/some", handler.GetSome)
	router.Post("/some", handler.CreateSome)

	// TODO add middleware

	return router
}
