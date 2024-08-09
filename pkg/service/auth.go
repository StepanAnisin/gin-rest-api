package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/StepanAnisin/gin-rest-api/models"
	"github.com/StepanAnisin/gin-rest-api/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const salt = "zdarovabratan12345"
const signKey = "123#$%FASD"
const tokenLT = time.Minute * 10

type AuthService struct {
	AuthRepo repository.AuthRepository
}

type tokenClaims struct {
	jwt.StandardClaims
	userName string `json:"user_name"`
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{AuthRepo: repo}
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.userName, nil
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.AuthRepo.CreateUser(user)
}

func (s *AuthService) GetUserById(id int) (*models.User, error) {
	return s.AuthRepo.GetUserById(id)
}

func (s *AuthService) GetUserByLogin(login string) (*models.User, error) {
	return s.AuthRepo.GetUserByLogin(login)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.AuthRepo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenLT).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Username,
	})

	return token.SignedString([]byte(signKey))
}
