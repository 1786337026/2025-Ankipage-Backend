package services

import (
	"Ankipage/db"
	"Ankipage/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Register(username, password, email string) (models.User, error) {

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("error hashing password")
	}

	// 创建用户
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return models.User{}, errors.New("error creating user")
	}
	return user, nil
}
