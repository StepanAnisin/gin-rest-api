package main

import (
	"github.com/StepanAnisin/gin-rest-api"
	"github.com/StepanAnisin/gin-rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running htt server %s", err.Error())
	}
}
