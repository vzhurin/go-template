package main

import (
	"errors"
	"github.com/vzhurin/template/internal/task/application"
	"github.com/vzhurin/template/internal/task/infrastructure/http/api"
	"log"
	"net/http"
)

func main() {
	app, cleanup := application.NewApplication()
	defer cleanup()
	server := api.NewServer(app)
	router := api.NewRouter(server)
	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}
	err := httpServer.ListenAndServe()

	// TODO add graceful shut down

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
