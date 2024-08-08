package repository

import (
	"github.com/StepanAnisin/gin-rest-api/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (int, error) {
	// Сохраняем пользователя в базе данных
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.Id, nil
}

func (r *AuthRepository) GetUser(username, password string) (*models.User, error) {
	var user models.User
	// Поиск пользователя по логину и паролю
	result := r.db.Where("login = ? AND password = ?", username, password).First(&user)

	// Проверяем, найден ли пользователь
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // пользовтель не найден
		}
		return nil, result.Error // ошибка при запросе
	}

	return &user, nil // возвращаем найденного пользователя

}
