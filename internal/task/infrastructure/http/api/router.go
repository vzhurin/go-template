package api

import (
	"net/http"
)

func NewRouter(server ServerInterface) http.Handler {
	router := HandlerWithOptions(server, ChiServerOptions{
		BaseURL: "/api",
	})

	// TODO add middlewares

	return router
}
