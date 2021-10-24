package main

import (
	"log"
	"os"

	"gitlab.com/zharzhanov/region"
	"gitlab.com/zharzhanov/region/database/mongo"
	"gitlab.com/zharzhanov/region/pkg/handler"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"gitlab.com/zharzhanov/region/pkg/service"
)

// @title Swagger Region.app Объявления API
// @version 1.0
// @description REST API for Region.app Объявления.
// @termsOfService http://swagger.io/terms/

// @host localhost
// @BasePath /
func main() {

	db := mongo.NewMongoDB(mongo.Config{
		MongoUser:     os.Getenv("mongo_user"),
		MongoPassword: os.Getenv("mongo_password"),
		MongoPort:     os.Getenv("mongo_port"),
		MongoHost:     os.Getenv("mongo_host"),
		DbName:        os.Getenv("mongo_db"),
	})
	//mongo.CreateIndexes(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(region.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured during starting web service: %s", err.Error())
	}
}
