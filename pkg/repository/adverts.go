package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdvertMongo struct {
	db *mongo.Database
}

func NewAdvertMongo(db *mongo.Database, advertCollection string) *AdvertMongo {
	return &AdvertMongo{
		db:db,
	}
}

func (r *AdvertMongo) CreateAdvert(ctx context.Context, advert models.AdvertInput) (string, error) {
	res, err := r.db.Collection(advertsCollection).InsertOne(ctx, advert)

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AdvertMongo) UploadImage(ctx context.Context, advertId string, urls []string) error {
	objId, _ := primitive.ObjectIDFromHex(advertId)

	objList := []interface{}{}
	for url := range urls {
		objList = append(objList, bson.M{"url": url})
	}

	res, err := r.db.Collection(imageCollection).InsertMany(ctx, objList)
	if err != nil {
		return err
	}

	_, err = r.db.Collection(advertsCollection).UpdateOne(ctx, bson.M{"_id": objId},  bson.M{"$push": bson.M{"images": bson.M{"$each": res.InsertedIDs}}})
	if err != nil {
		return err
	}

	return err
}

func (r *AdvertMongo) GetAllAdverts(ctx context.Context, filter bson.M) ([]models.Advert, error) {
	adverts := make([]models.Advert, 0)

	cur, err := r.db.Collection(advertsCollection).Find(ctx, filter)
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
	err := r.db.Collection(advertsCollection).FindOne(ctx,bson.M{"_id": objId}).Decode(&advert)

	if err != nil {
		return nil, err
	}

	return models.ToAdvert(advert), nil
}
func (r *AdvertMongo) UpdateAdvert(ctx context.Context, id string, advert models.UpdateAdvertInput) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.Collection(advertsCollection).UpdateOne(ctx, bson.D{{"_id", objId}},  bson.M{ "$set": advert})
	if err != nil {
		return err
	}

	return nil
}

func (r *AdvertMongo) DeleteAdvert(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := r.db.Collection(advertsCollection).DeleteOne(ctx, bson.D{{"_id", objId}})
	if err != nil {
		return err
	}

	return nil
}
