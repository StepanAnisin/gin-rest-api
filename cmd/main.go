package main

import (
	"context"
	"github.com/StepanAnisin/gin-rest-api/internal"
	"github.com/StepanAnisin/gin-rest-api/internal/config"
	"gorm.io/driver/postgres"

	"github.com/StepanAnisin/gin-rest-api/pkg/handler"
	"github.com/StepanAnisin/gin-rest-api/pkg/repository"
	"github.com/StepanAnisin/gin-rest-api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

// @Title gin-template API
// @version 1.0
// @description API server for template Application

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	logrus.SetFormatter(new(logrus.JSONFormatter))
	conf := config.New()
	dbConnection, err := connect(conf.Database.ConnectionString)
	if err != nil {
		logrus.Fatal("database connection error")
	}
	authRepo := repository.NewAuthRepository(dbConnection)
	authService := service.NewAuthorizationService(authRepo)
	handlers := handler.NewHandler(authService)
	srv := internal.NewHttpServer(conf, handlers)

	srv.Run(ctx)
	logrus.Print("App Started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sign := <-quit:
		logrus.Debug("received %s signal, server shutting down", sign)
		cancel()
	case <-ctx.Done():
		logrus.Debug("received cancel signal, server shutting down")
	}

	logrus.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
