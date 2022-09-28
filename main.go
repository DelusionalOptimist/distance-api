package main

import (
	"log"
	"os"

	"github.com/DelusionalOptimist/distance-api/router"
)

func main() {
	logger := log.New(os.Stdout, "distance-api: ", 0)
	router := router.NewRouter(8080, logger)

	logger.Println("Starting listening on localhost:8080")
	if err := router.Run(); err != nil {
		logger.Fatalln(err)
	}
}
