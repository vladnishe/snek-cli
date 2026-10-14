package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"os"

	"github.com/vladnishe/snek-cli/internal/config"
	"github.com/vladnishe/snek-cli/internal/gemini"
	"github.com/vladnishe/snek-cli/internal/weather"
	"github.com/vladnishe/snek-cli/internal/weather/validator"
)

const (
	devEnv  = "dev"
	prodEnv = "prod"
)

func main() {
	var city string
	flag.StringVar(&city, "city", "", "city for weather forecast")
	flag.Parse()

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	if err := validator.ValidateCity(city); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Debug("initializing gemini and weather clients...")
	_ = gemini.New(cfg.GeminiAPIKey, cfg.GeminiAPIBaseURL)
	wClient := weather.New(cfg.WeatherAPIKey, cfg.WeatherAPIBaseURL)

	fmt.Printf("Requesting weather forecast for: %s....\n", city)
	forecast, err := wClient.GetWeather(context.Background(), city)
	if err != nil {
		log.Error("failed to get forecast: ", "err", err)
		os.Exit(1)
	}

	fmt.Printf("The weather in %s: %.0f, %s\n", city, math.Floor(forecast.Main.Temp), forecast.Weather[0].Description)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case devEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case prodEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return log
}
