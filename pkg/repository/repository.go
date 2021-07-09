package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authentication interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, user models.User) (string, error)
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.Advert) (string, error)
	GetAllAdverts(ctx context.Context) ([]models.Advert, error)
	GetAdvertById(ctx context.Context, id string) (*models.Advert, error)
	UpdateAdvert(ctx context.Context, id string,  advert models.UpdateAdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
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
		Adverts : NewAdvertMongo(db, "adverts"),
	}
}
