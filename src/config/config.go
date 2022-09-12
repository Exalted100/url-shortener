package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v9"
)

type ConfigVariablesType struct {
	RedisHost string `validate:"required"`
	RedisPort string `validate:"required"`
	Port      string `validate:"required"`
	AppHost   string `validate:"required"`
}

var ConfigVariables ConfigVariablesType

func SetConfig() {
	if os.Getenv("APP_ENV") != "prod" || os.Getenv("APP_ENV") != "stg" || os.Getenv("APP_ENV") != "beta" {
		godotenv.Load("./env/dev.env")
	}

	ConfigVariables = ConfigVariablesType{
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		Port:      os.Getenv("PORT"),
		AppHost:   os.Getenv("APP_HOST"),
	}

	if ConfigVariables.RedisHost == "" {
		ConfigVariables.RedisHost = "localhost"
	}

	if ConfigVariables.RedisPort == "" {
		ConfigVariables.RedisPort = "6379"
	}

	if ConfigVariables.Port == "" {
		ConfigVariables.Port = "8000"
	}

	if ConfigVariables.AppHost == "" {
		ConfigVariables.AppHost = "localhost:8000"
	}

	validate := validator.New()
	err := validate.Struct(ConfigVariables)

	if err != nil {
		log.Fatalf("Environmental variables have not been set: %v", err)
	}

	return
}
