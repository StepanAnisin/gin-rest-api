package service

import (
	"github.com/StepanAnisin/gin-rest-api/models"
	"github.com/StepanAnisin/gin-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
