package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo *repository.Repository) *AdminService {
	return &AdminService{repo: repo.Admin}
}

func (s *AdminService) GetUser(ctx context.Context, phone string) (string, error) {

	id, _ := primitive.ObjectIDFromHex("61695a227f5f834c675c07fd")
	userID := id.Hex()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(15 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	generatedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return generatedToken, err
}