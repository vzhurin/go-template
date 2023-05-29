package api

import (
	"github.com/rs/cors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	common "github.com/vzhurin/template/internal/common/infrastructure/http"
)

func NewHandler(server ServerInterface, logger logrus.FieldLogger) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(common.NewLoggingMiddleware(logger))
	router.Use(cors.Default().Handler)

	return HandlerWithOptions(server, ChiServerOptions{
		BaseURL:    "/api",
		BaseRouter: router,
	})
}
