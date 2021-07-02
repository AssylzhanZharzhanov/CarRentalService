package main

import (
	"gitlab.com/zharzhanov/region"
	"log"
	"os"
)

func main() {
	srv := new(region.Server)
	log.Println(os.Getenv("PORT"))
	srv.Run(os.Getenv("PORT"))
}
