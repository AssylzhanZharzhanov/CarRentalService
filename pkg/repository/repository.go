package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authentication interface {
	SignUp(ctx context.Context, user models.User) (string, error)
	SignIn(ctx context.Context, user models.User) (string, error)
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.Advert) (string, error)
	GetAllAdverts(ctx context.Context) error
}

type Users interface {

}

type Repository struct {
	Authentication
	Adverts
	Users
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
	}
}
