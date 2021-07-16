package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateIndexes(db *mongo.Database) {
	if err := PopulateIndex(db, "adverts"); err != nil {
		log.Printf("%s", err.Error())
	}
}

func PopulateIndex(db *mongo.Database, collection string) error {
	adverts := mongo.IndexModel{
		Keys: bson.M{
			"title_search": 1,
		},
	}
	_, err := db.Collection(collection).Indexes().CreateOne(context.Background(), adverts)
	if err != nil {
		return err
	}

	return nil
}
