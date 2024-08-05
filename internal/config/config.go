package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Data1URL  string
	Data2URL  string
	CreateURL string
	Port      string
}

func LoadConfig() *Config {
	cfg := &Config{
		Data1URL:  getEnv("DATA1_URL", "https://viacep.com.br/ws/"),
		Data2URL:  getEnv("DATA2_URL", "https://pokeapi.co/api/v2/pokemon/ditto"),
		CreateURL: getEnv("DATA1_URL", "https://default-data1-url.com"),
		Port:      getEnv("PORT", ":9091"),
	}

	log.Info("Configuration loaded successfully")

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
