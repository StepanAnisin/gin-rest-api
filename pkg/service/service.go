package service

import (
	"github.com/StepanAnisin/gin-rest-api/models"
	"github.com/StepanAnisin/gin-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (string, error)
	GetUserById(id int) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
}

type AuthorizationService struct {
	Authorization
}

func NewAuthorizationService(repo *repository.AuthRepository) *AuthorizationService {
	return &AuthorizationService{
		Authorization: NewAuthService(*repo),
	}
}
