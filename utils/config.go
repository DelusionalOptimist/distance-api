package utils

import "os"

// contains the configuration for the server
type Config struct {
	Port   string
	APIKey string
}

// GlobalConfig is used for storing API keys and
// other useful configuration
var GlobalConfig Config

// initializes the GlobalConfig
func NewConfig() error {
	apiKey := os.Getenv("API_KEY")
	port := os.Getenv("PORT")

	GlobalConfig = Config{
		Port:   port,
		APIKey: apiKey,
	}

	return nil
}
