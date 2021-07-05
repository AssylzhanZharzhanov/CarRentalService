package main

import (
	"github.com/spf13/viper"
	"gitlab.com/zharzhanov/region"
	"gitlab.com/zharzhanov/region/pkg/handler"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"gitlab.com/zharzhanov/region/pkg/service"
	"log"
)

func main() {

	//if err := initConfig(); err != nil {
	//	log.Fatalf("Error initializing server configs: %s", err.Error())
	//}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(region.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured during starting web service: %s", err.Error())
	}
}

func initConfig () error {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath("/configs")
	return viper.ReadInConfig()
}