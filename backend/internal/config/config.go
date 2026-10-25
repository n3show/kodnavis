package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env         string
	Port        string
	DatabaseURL string
}

func Load() (Config, error) {
	cfg := Config{
		Env:         getenv("APP_ENV", "development"),
		Port:        getenv("API_PORT", "8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	return cfg, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
