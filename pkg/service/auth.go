package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repository *repository.Repository) *AuthService {
	return &AuthService{
		repo: repository.Authentication,
	}
}

func (s *AuthService) SignIn(ctx context.Context, user models.User) (string, error) {
	return s.repo.SignIn(ctx, user)
}

func (s *AuthService) SignUp(ctx context.Context, user models.User) (string, error) {
	return s.repo.SignUp(ctx, user)
}

func generatePassword(password string)  string {
	return ""
}