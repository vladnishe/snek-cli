package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	DefaultWeatherAPIBaseURL = "https://api.openweathermap.org/data/2.5"
	DefaultGroqAPIBaseURL    = "https://api.groq.com/openai/v1"
)

type Config struct {
	Env               string `env:"ENV"`
	GROQAPIKey        string `env:"GROQ_API_KEY"`
	GroqModel         string `env:"GROQ_MODEL"`
	WeatherAPIKey     string `env:"WEATHER_API_KEY"`
	WeatherAPIBaseURL string `env:"WEATHER_BASE_URL"`
	GROQAPIBaseURL    string `env:"GROQ_BASE_URL"`
}

func MustLoad() *Config {
	_ = godotenv.Load()

	return &Config{
		Env:               getEnv("ENV", "prod"),
		GROQAPIKey:        getApiEnv("GROQ_API_KEY"),
		GroqModel:         getEnv("GROQ_MODEL", "openai/gpt-oss-120b"),
		WeatherAPIKey:     getApiEnv("WEATHER_API_KEY"),
		WeatherAPIBaseURL: getEnv("WEATHER_BASE_URL", DefaultWeatherAPIBaseURL),
		GROQAPIBaseURL:    getEnv("GROQ_BASE_URL", DefaultGroqAPIBaseURL),
	}
}

func getEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func getApiEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("%s env is not specified", key))
	}
	return val
}
