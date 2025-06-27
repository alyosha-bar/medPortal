package services

import (
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
)

func SignUp(user models.User) error {
	return repository.SignUp(user)
}

func GetUserByUsername(username string) (models.User, error) {
	return repository.GetUserByUsername(username)
}
