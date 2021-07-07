package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{
		db: db,
	}
}

func (r *AuthMongo) SignUp(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

func (r *AuthMongo) SignIn(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

