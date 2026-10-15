package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"os"

	"github.com/vladnishe/snek-cli/internal/ai"
	"github.com/vladnishe/snek-cli/internal/config"
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
	askAdvice := flag.Bool("advice", false, "ask advice for a day")
	flag.Parse()

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	if err := validator.ValidateCity(city); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Debug("initializing gemini and weather clients...")
	aiClient := ai.New(cfg.GROQAPIKey, cfg.GROQAPIBaseURL, cfg.GroqModel)
	wClient := weather.New(cfg.WeatherAPIKey, cfg.WeatherAPIBaseURL)

	fmt.Printf("Requesting weather forecast for: %s....\n", city)
	forecast, err := wClient.GetWeather(context.Background(), city)
	if err != nil {
		log.Error("failed to get forecast: ", "err", err)
		os.Exit(1)
	}

	if *askAdvice {
		advice, err := aiClient.GetAdviceForWeather(
			context.Background(),
			city,
			forecast.Weather[0].Description,
			forecast.Main.Temp,
			forecast.Wind.Speed,
			forecast.Rain.Precipation,
			forecast.Clouds.Percentage,
		)
		if err != nil {
			log.Error("failed to get advice: ", "err", err)
		}
		fmt.Println(advice)
	} else {
		fmt.Printf("The weather in %s: %.0f, %s\n", city, math.Floor(forecast.Main.Temp), forecast.Weather[0].Description)
	}
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
