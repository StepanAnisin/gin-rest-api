package main

import (
	"github.com/StepanAnisin/gin-rest-api"
	"github.com/StepanAnisin/gin-rest-api/pkg/handler"
	"github.com/StepanAnisin/gin-rest-api/pkg/repository"
	"github.com/StepanAnisin/gin-rest-api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running htt server %s", err.Error())
	}
}
