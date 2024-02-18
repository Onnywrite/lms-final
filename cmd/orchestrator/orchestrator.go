package main

import (
	"github.com/Onnywrite/lms-final/internal/app"
	"github.com/Onnywrite/lms-final/internal/config"
	"github.com/Onnywrite/lms-final/internal/domain"
	"log/slog"
)

func main() {
	const op = "main"

	cfg := config.MustLoadOrchestrator()

	logger := domain.MustSetupLoggerInDir(cfg.Env, cfg.LogsDir)
	log := logger.With(slog.String("op", op))

	application := app.NewOrchestrator(logger, cfg.Port, cfg.DbConnection, cfg.AllowOrigin)
	log.Debug("Starting the orchestrator")
	if err := application.Start(); err != nil {
		log.Error("Could not start the orchestrator")
		return
	}

	// graceful shutdown
	domain.ShutdownOnSignal(cfg.ShutdownTimeout, logger, application)
}
