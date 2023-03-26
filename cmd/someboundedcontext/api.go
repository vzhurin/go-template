package main

import (
	"errors"
	"log"
	nethttp "net/http"

	"template/internal/someboundedcontext/application"
	"template/internal/someboundedcontext/infrastructure/http"
)

func main() {
	app, cleanup := application.NewApplication()
	defer cleanup()
	handler := http.NewHandler(app)
	router := http.NewRouter(handler)
	server := http.NewHTTPServer(":5000", router)
	err := server.ListenAndServe()

	if errors.Is(err, nethttp.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
