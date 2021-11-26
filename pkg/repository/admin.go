package repository

import "go.mongodb.org/mongo-driver/mongo"

type AdminMongo struct {
	db *mongo.Database
}

func NewAdminMongo(db *mongo.Database) *AdminMongo {
	return &AdminMongo{
		db: db,
	}
}

