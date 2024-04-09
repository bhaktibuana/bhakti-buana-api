package configs

import (
	"os"
)

type S_DBConfig struct {
	DB_CONNECTION string
	DB_DSN        string
	DB_DATABASE   string
}

func DBConfig() S_DBConfig {
	return S_DBConfig{
		DB_CONNECTION: os.Getenv("DB_CONNECTION"),
		DB_DSN:        os.Getenv("DB_DSN"),
		DB_DATABASE:   os.Getenv("DB_DATABASE"),
	}
}
