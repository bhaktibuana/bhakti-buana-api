package configs

import (
	"os"
)

type S_AppConfig struct {
	PORT           string
	GIN_MODE       string
	BASE_URL       string
	JWT_SECRET_KEY string
}

func AppConfig() S_AppConfig {
	return S_AppConfig{
		PORT:           ":" + os.Getenv("PORT"),
		GIN_MODE:       os.Getenv("GIN_MODE"),
		BASE_URL:       os.Getenv("BASE_URL"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
	}
}
