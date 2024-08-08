package models

type User struct {
	Id       int    `json:"-" gorm:"primaryKey" sql:"AUTO_INCREMENT"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
