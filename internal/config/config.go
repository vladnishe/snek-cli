package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	DefaultWeatherAPIBaseURL = "https://api.openweathermap.org/data/2.5"
	DefaultGeminiAPIBaseURL  = "https://generativelanguage.googleapis.com/v1beta"
)

type Config struct {
	Env               string `env:"ENV"`
	GeminiAPIKey      string `env:"GEMINI_API_KEY"`
	WeatherAPIKey     string `env:"WEATHER_API_KEY"`
	WeatherAPIBaseURL string `env:"WEATHER_BASE_URL"`
	GeminiAPIBaseURL  string `env:"GEMINI_BASE_URL"`
}

func MustLoad() *Config {
	_ = godotenv.Load()

	return &Config{
		Env:               getEnv("ENV", "prod"),
		GeminiAPIKey:      getApiEnv("GEMINI_API_KEY"),
		WeatherAPIKey:     getApiEnv("WEATHER_API_KEY"),
		WeatherAPIBaseURL: getEnv("WEATHER_BASE_URL", DefaultWeatherAPIBaseURL),
		GeminiAPIBaseURL:  getEnv("GEMINI_BASE_URL", DefaultGeminiAPIBaseURL),
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
