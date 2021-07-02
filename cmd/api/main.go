package main

import (
	"gitlab.com/zharzhanov/region"
	"gitlab.com/zharzhanov/region/pkg/handler"
	"log"
	"os"
)

func main() {
	handler := new(handler.Handler)
	srv := new(region.Server)
	if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured during starting web service: %s", err.Error())
	}
}
