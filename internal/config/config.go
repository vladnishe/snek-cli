package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string `env:"ENV"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic("failed to parse env vars for config initializing")
	}

	return &Config{
		Env: os.Getenv("ENV"),
	}
}
