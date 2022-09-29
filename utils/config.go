package utils

import "os"

type Config struct {
	Port string
	APIKey string
}

var GlobalConfig Config

func NewConfig() error {
	apiKey := os.Getenv("API_KEY")
	port := os.Getenv("PORT")

	GlobalConfig = Config{
		Port: port,
		APIKey: apiKey,
	}

	return nil
}
