package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	codeCollection = "codes"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{
		db: db,
	}
}

func (r *AuthMongo) VerifyCode(ctx context.Context, code string) error {
	//var code models.Code
	err := r.db.Collection(codeCollection).FindOne(ctx, bson.M{"code": code})
	log.Println(err)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (r *AuthMongo) CreateUser(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

func (r *AuthMongo) GetUser(ctx context.Context, user models.User) (string, error) {
	return "", nil
}

func (r *AuthMongo) CreateCode(ctx context.Context, code models.Code) error {
	_, err := r.db.Collection(codeCollection).InsertOne(ctx, code)
	if err != nil {
		return err
	}
	return err
}


