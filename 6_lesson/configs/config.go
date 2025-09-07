package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Db   string
	Temp string
}

func LoadConfig() *AppConfig {

	err := godotenv.Load()

	if err != nil {

	}

	return &AppConfig{
		Db:   "example",
		Temp: os.Getenv("EXAMPLE"),
	}
}
