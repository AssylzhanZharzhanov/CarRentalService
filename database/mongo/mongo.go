package mongo

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type Config struct {
	MongoUser string
	MongoPassword string
	MongoHost string
	MongoPort string
	DbName string
}

const (
	usersCollection = "users"
	advertsCollection = "adverts"
	paymentsCollection = "payments"
	carBrandsCollection = "car_brands"
)

func NewMongoDB(cfg Config) *mongo.Database {
	MongoURI :=  fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		cfg.MongoUser,
		cfg.MongoPassword,
		cfg.MongoHost,
		cfg.MongoPort,
	)

	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		logrus.Fatalf("%s", err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		logrus.Printf("%s", err.Error())
	} else {
		logrus.Println("Connected to DB")
	}

	return client.Database(cfg.DbName)
}

