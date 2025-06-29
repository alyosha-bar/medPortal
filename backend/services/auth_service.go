package services

import (
	"github.com/alyosha-bar/medPortal/models"
	"github.com/alyosha-bar/medPortal/repository"
)

type AuthService struct {
	Repo *repository.AuthRepo
}

func NewAuthService(repo *repository.AuthRepo) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) SignUp(user models.User) error {
	return s.Repo.SignUp(user)
}

func (s *AuthService) GetUserByUsername(username string) (models.User, error) {
	return s.Repo.GetUserByUsername(username)
}
