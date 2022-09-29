package main

import (
	"log"
	"os"

	"github.com/DelusionalOptimist/distance-api/router"
	"github.com/DelusionalOptimist/distance-api/utils"
)

func main() {
	logger := log.New(os.Stdout, "distance-api: ", 0)

	if err := utils.NewConfig(); err != nil {
		logger.Fatalln(err)
	}

	router := router.NewRouter(logger)

	logger.Println("Starting listening on localhost:8080")
	if err := router.Run(); err != nil {
		logger.Fatalln(err)
	}
}
