package main

import (
	"log"

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
		MongoUser:     "mongo",
		MongoPassword: "mongo",
		MongoPort:     "27017",
		MongoHost:     "mongo",
		DbName:        "region",
	})
	mongo.CreateIndexes(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(region.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured during starting web service: %s", err.Error())
	}
}
