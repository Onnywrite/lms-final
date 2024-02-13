package main

import (
	"context"
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
	cfg := config.MustLoadOrchestrator()

	log := setupLogger(cfg.Env)

	application := app.NewOrchestrator(log, cfg.Port, cfg.DbConnection)
	log.Debug("Starting the orchestrator")
	go application.MustStart()

	// graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)
	log.Debug("Waiting for shutdown in orchestrator.go")

	<-shutdown
	log.Info("Shutting down")

	c, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()
	application.Stop(c)
	select {
	case <-c.Done():
		log.Error("Shutdown timeout exceeded")
	}
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