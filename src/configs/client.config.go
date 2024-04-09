package configs

import (
	"os"
)

type S_ClientConfig struct {
	CLIENT_URL string
}

func ClientConfig() S_ClientConfig {
	return S_ClientConfig{
		CLIENT_URL: os.Getenv("CLIENT_URL"),
	}
}
