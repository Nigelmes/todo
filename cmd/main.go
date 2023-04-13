package main

import (
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/pkg/handler"
	"github.com/nigelmes/todo/pkg/repository"
	"github.com/nigelmes/todo/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err)
	}
}
