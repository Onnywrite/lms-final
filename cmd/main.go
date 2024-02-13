package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Onnywrite/lms-final/internal/app"
	"github.com/Onnywrite/lms-final/internal/config"
)

const (
	localEnv = "loc"
	devEnv   = "dev"
	prodEnv  = "prod"
)

func main() {
	cfg := config.MustLoadMain()

	log := setupLogger(cfg.Env)
	defer log.Debug("stopped")
	application := app.NewMain(log, cfg.Port, cfg.StaticDir)
	log.Debug("Starting the application")
	go application.MustStart()

	// graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	log.Debug("Waiting for shutdown")
	<-shutdown
	application.Stop()
	log.Info("Gracefully stopped")
}

func setupLogger(env string) *slog.Logger {
	switch env {
	case localEnv:
		return slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case devEnv:
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prodEnv:
		return slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		panic("unexpected env value")
	}
}
