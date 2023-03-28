package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/vzhurin/template/internal/task/application"
	"github.com/vzhurin/template/internal/task/infrastructure/http/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	AppName = "template"
)

func main() {
	l := logrus.StandardLogger()
	l.SetFormatter(&logrus.JSONFormatter{})
	logger := l.WithField("app", AppName)

	app, cleanup := application.NewApplication(logger)
	defer cleanup()
	server := api.NewServer(app)
	router := api.NewRouter(server)
	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to start HTTP server: %s\n", err)
		}
	}()
	logger.Println("HTTP server started")

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatalf("failed to gracefully shutdown HTTP server: %s\n", err)
	}
	logger.Println("HTTP server was gracefully shut down")
}
