package main

import (
	"api/internal"
	"api/internal/handlers"
	"api/internal/repository"
	"api/internal/service"
	"api/pkg/database"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {

	db, err := database.InitDb()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	srv := new(internal.Server)
	go func() {
		if err := srv.Run("8080", todoHandler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
