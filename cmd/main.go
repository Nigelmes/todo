package main

import (
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/config"
	"github.com/nigelmes/todo/pkg/handler"
	"github.com/nigelmes/todo/pkg/repository"
	"github.com/nigelmes/todo/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	cfg := config.GetConfig()

	server := new(todo.Server)
	if err := server.Run(cfg, handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err)
	}
}
