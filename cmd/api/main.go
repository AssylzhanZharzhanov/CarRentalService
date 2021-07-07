package main

import (
	"gitlab.com/zharzhanov/region"
	"gitlab.com/zharzhanov/region/pkg/handler"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"gitlab.com/zharzhanov/region/pkg/service"
	"log"
)

func main() {

	db := repository.NewMongoDB(repository.Config{
		MongoUser: "mongo",
		MongoPassword: "mongo",
		MongoPort: "27017",
		MongoHost: "mongo",
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(region.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured during starting web service: %s", err.Error())
	}
}
