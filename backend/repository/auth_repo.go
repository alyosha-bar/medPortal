package repository

import (
	"github.com/alyosha-bar/medPortal/models"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (r *AuthRepo) SignUp(user models.User) error {
	result := r.DB.Create(&user)
	return result.Error
}

func (r *AuthRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
