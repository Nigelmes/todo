package main

import (
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/config"
	"github.com/nigelmes/todo/pkg/handler"
	"github.com/nigelmes/todo/pkg/repository"
	"github.com/nigelmes/todo/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("error connection db, %s", err)
	}
	logrus.Println("DB connection successful")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(cfg, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err)
	}
}
