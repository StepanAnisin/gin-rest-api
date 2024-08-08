package handler

import (
	"github.com/StepanAnisin/gin-rest-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *service.AuthorizationService
}

func NewHandler(authService *service.AuthorizationService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		//Вызов метода userIdentity для дальнейшей реализации api
		//api := router.Group("/api", h.userIdentity)
	}
	return router
}
