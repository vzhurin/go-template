package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: "localhost:5000",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			io.WriteString(writer, "Hello World!")
		}),
	}
	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
