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
	cfg := config.GetConfig()

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("error connection db, %s", err)
	}
	log.Println("DB connection successful")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(cfg, handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err)
	}
}
