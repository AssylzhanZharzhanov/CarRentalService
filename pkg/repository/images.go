package repository

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImageMongo struct {
	db *mongo.Database
}

func (r *ImageMongo) DeleteImage(ctx context.Context, imageId string, advertId string) error {
	advertObjId, _ := primitive.ObjectIDFromHex(advertId)
	imageObjId, _ := primitive.ObjectIDFromHex(imageId)

	_, err := r.db.Collection(advertsCollection).UpdateOne(ctx, bson.M{"_id": advertObjId}, bson.M{"$pull": bson.M{"images": bson.M{"_id": imageObjId}}})
	if err != nil {
		return err
	}

	return nil
}

func (r *ImageMongo) UploadImage(ctx context.Context, advertId string, url string) error {
	objId, _ := primitive.ObjectIDFromHex(advertId)

	image := models.Image{
		ID: primitive.NewObjectID(),
		Url: url,
	}

	//res, err := r.db.Collection(imageCollection).InsertOne(ctx, image)
	//if err != nil {
	//	return err
	//}

	_, err := r.db.Collection(advertsCollection).UpdateOne(ctx, bson.M{"_id": objId},  bson.M{"$push": bson.M{"images": image}})
	if err != nil {
		return err
	}

	return nil
}

func (r *ImageMongo) GetImageById(ctx context.Context, id string) (models.Image, error) {
	objId, _ := primitive.ObjectIDFromHex(id)

	var image models.Image

	err := r.db.Collection(imageCollection).FindOne(ctx, bson.M{"_id":objId }).Decode(&image)

	if err != nil {
		return image, err
	}

	return image, nil
}

func NewImageMongo(db *mongo.Database) *ImageMongo {
	return &ImageMongo{db: db}
}
