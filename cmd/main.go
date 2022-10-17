package main

import (
	"github.com/StepanAnisin/gin-rest-api"
	"github.com/StepanAnisin/gin-rest-api/pkg/handler"
	"github.com/StepanAnisin/gin-rest-api/pkg/repository"
	"github.com/StepanAnisin/gin-rest-api/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running htt server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
