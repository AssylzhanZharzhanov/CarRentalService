package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	usersCollection = "users"
	codeCollection = "codes"
)

type AuthMongo struct {
	db *mongo.Database
}

func (r *AuthMongo) VerifyCode(ctx context.Context, code string) (models.Code, error) {
	var output models.Code
	err := r.db.Collection(codeCollection).FindOne(ctx, bson.M{"code": code}).Decode(&output)
	if err != nil {
		return output, err
	}
	return output, err
}

func (r *AuthMongo) CreateUser(ctx context.Context, user models.User) (string, error) {
	//	Временно
	res, err := r.db.Collection(usersCollection).InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AuthMongo) GetUser(ctx context.Context, phone string) (string, error) {
	user := models.User{}
	err := r.db.Collection(usersCollection).FindOne(ctx, bson.M{"phone": phone}).Decode(&user)

	return user.ID.Hex(), err
}

func (r *AuthMongo) CreateCode(ctx context.Context, code models.Code) error {
	_, err := r.db.Collection(codeCollection).InsertOne(ctx, code)
	if err != nil {
		return err
	}
	return err
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{
		db: db,
	}
}

