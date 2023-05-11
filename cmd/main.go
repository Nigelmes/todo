package main

import (
	"context"
	"github.com/nigelmes/todo"
	"github.com/nigelmes/todo/config"
	"github.com/nigelmes/todo/pkg/handler"
	"github.com/nigelmes/todo/pkg/repository"
	"github.com/nigelmes/todo/pkg/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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
	go func(){
		if err := server.Run(cfg, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error running http server: %s", err)
		}
	}()
	logrus.Print("app started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("app stopped")

	if err := server.ShutDown(context.Background()); err != nil{
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil{
		logrus.Errorf("error occured on db connection close : %s", err.Error())
	}
}
