package main

import (
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(todo.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err)
	}
}
