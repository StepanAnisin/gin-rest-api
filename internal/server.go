package internal

import (
	"context"
	"errors"
	"github.com/StepanAnisin/gin-rest-api/internal/config"
	"github.com/StepanAnisin/gin-rest-api/pkg/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxHeaderBytes = 1 << 20
)

type HttpServer struct {
	httpServer http.Server
}

func (s *HttpServer) Run(ctx context.Context) {
	go func() {
		_, cancel := context.WithCancel(ctx)
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Error("error occurred while running http server")
			cancel()
		}
	}()

	logrus.Debug("connect to database")
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewHttpServer(config *config.Config, handlers *handler.Handler) *HttpServer {
	return &HttpServer{
		httpServer: http.Server{
			Addr:           ":" + config.HttpConfig.Port,
			Handler:        handlers.InitRoutes(config.HttpConfig.Port),
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
		},
	}
}
