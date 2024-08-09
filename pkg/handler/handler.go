package handler

import (
	_ "github.com/StepanAnisin/gin-rest-api/cmd/docs"
	"github.com/StepanAnisin/gin-rest-api/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	authService *service.AuthorizationService
}

func NewHandler(authService *service.AuthorizationService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		//Вызов метода userIdentity для дальнейшей реализации api
		//api := router.Group("/api", h.userIdentity)
	}
	return router
}
