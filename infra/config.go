package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Environment  string
	App          string
	AccessSecret string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	return Config{
		Environment:  os.Getenv("ENVIRONMENT"),
		App:          os.Getenv("APP"),
		AccessSecret: os.Getenv("ACCESS_SECRET"),
	}
}
