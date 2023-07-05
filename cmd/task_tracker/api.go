package main

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vzhurin/template/internal/task_tracker/application"
	"github.com/vzhurin/template/internal/task_tracker/infrastructure/http/api"
	"github.com/vzhurin/template/internal/task_tracker/infrastructure/persistance"
)

const (
	AppName = "task_tracker_api"
)

func main() {
	l := logrus.StandardLogger()
	l.SetFormatter(&logrus.JSONFormatter{})
	logger := l.WithField("app", AppName)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to database")
	}

	err = db.AutoMigrate(&persistance.GormTask{})
	if err != nil {
		logger.Fatalf("failed to migrate")
	}

	taskRepository := persistance.NewGormTaskRepository(db)
	// TODO
	taskFinder := persistance.NewInMemoryTaskFinder()
	app, cleanup := application.NewApplication(taskRepository, taskFinder, logger)
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
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Fatalf("failed to gracefully shutdown HTTP server: %s\n", err)
	}
	logger.Println("HTTP server was gracefully shut down")
}
