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

func (r *AuthMongo) CreateUser(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

func (r *AuthMongo) GetUser(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

