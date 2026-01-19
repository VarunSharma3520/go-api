package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MongoURI string
	Env      string
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
    log.Println(".env file not found, using environment variables")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGODB_URI"),
		Env:      os.Getenv("ENV"),
	}

	// Optional: warn if something critical is missing
	if cfg.MongoURI == "" {
		log.Println("⚠️ Warning: MONGODB_URI is not set — using default localhost URI.")
	}

	return cfg
}
