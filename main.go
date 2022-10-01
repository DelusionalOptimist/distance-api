package main

import (
	"log"
	"os"

	"github.com/DelusionalOptimist/distance-api/router"
	"github.com/DelusionalOptimist/distance-api/utils"
)

func main() {
	// create a new global logger
	logger := log.New(os.Stdout, "distance-api ", 3)

	// initialize config from env variables
	// this config contains info like Google API key
	if err := utils.NewConfig(); err != nil {
		logger.Fatalln(err)
	}

	// create a router object and run
	router := router.NewRouter(logger)

	logger.Printf("Starting listening on localhost:%s", utils.GlobalConfig.Port)
	if err := router.Run(); err != nil {
		logger.Fatalln(err)
	}
}
