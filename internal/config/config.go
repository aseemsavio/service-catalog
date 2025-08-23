package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env      string
	HTTPPort string
	PG       PG
	LogLevel string
}

type PG struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	SSLMode  string
}

func Load() Config {
	return Config{
		Env:      get("APP_ENV", "dev"),
		HTTPPort: get("HTTP_PORT", "8080"),
		PG: PG{
			Host:     get("PG_HOST", "localhost"),
			Port:     get("PG_PORT", "5432"),
			User:     get("PG_USER", "app"),
			Password: get("PG_PASSWORD", "app"),
			DB:       get("PG_DB", "app"),
			SSLMode:  get("PG_SSLMODE", "disable"),
		},
		LogLevel: get("LOG_LEVEL", "info"),
	}
}

func (c Config) PGDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.PG.Host, c.PG.Port, c.PG.User, c.PG.Password, c.PG.DB, c.PG.SSLMode,
	)
}

func get(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
