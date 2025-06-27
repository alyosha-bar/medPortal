package repository

import (
	"github.com/alyosha-bar/medPortal/database"
	"github.com/alyosha-bar/medPortal/models"
)

func SignUp(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
