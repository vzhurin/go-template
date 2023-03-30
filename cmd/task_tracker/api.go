package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/vzhurin/template/internal/task_tracker/application"
	"github.com/vzhurin/template/internal/task_tracker/infrastructure/http/api"
	"github.com/vzhurin/template/internal/task_tracker/infrastructure/persistance"
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

	taskRepository := persistance.NewInMemoryTaskRepository()
	app, cleanup := application.NewApplication(taskRepository, taskRepository, logger)
	defer cleanup()
	server := api.NewServer(app, logger)
	handler := api.NewHandler(server, logger)
	httpServer := &http.Server{
		Addr:    ":5000",
		Handler: handler,
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
