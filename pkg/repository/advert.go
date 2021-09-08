package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdvertMongo struct {
	db *mongo.Collection
}

func NewAdvertMongo(db *mongo.Database, collection string) *AdvertMongo {
	return &AdvertMongo{db:db.Collection(collection)}
}

func (r *AdvertMongo) CreateAdvert(ctx context.Context, advert models.Advert) (string, error) {
	res, err := r.db.InsertOne(ctx, advert)

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AdvertMongo) GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error) {
	var adverts []models.Advert

	cur, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &adverts); err != nil {
		return nil, err
	}

	return adverts, nil
}

func (r *AdvertMongo) GetAdvertById(ctx context.Context, id string) (*models.Advert, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	var advert *models.Advert
	err := r.db.FindOne(ctx,bson.M{"_id": objId}).Decode(&advert)

	if err != nil {
		return nil, err
	}

	return models.ToAdvert(advert), nil
}
func (r *AdvertMongo) UpdateAdvert(ctx context.Context, id string, advert models.UpdateAdvertInput) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.UpdateOne(ctx, bson.D{{"_id", objId}},  bson.M{ "$set": advert})
	if err != nil {
		return err
	}

	return nil
}

func (r *AdvertMongo) DeleteAdvert(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.DeleteOne(ctx, bson.D{{"_id", objId}})
	if err != nil {
		return err
	}

	return nil
}
