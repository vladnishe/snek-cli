package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/vladnishe/snek-cli/internal/config"
	"github.com/vladnishe/snek-cli/internal/weather/util"
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

	if err := util.ValidateCity(city); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Requesting weather forecast for: %s....\n", city)
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
