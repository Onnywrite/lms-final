package main

import (
	"github.com/Onnywrite/lms-final/internal/app"
	"github.com/Onnywrite/lms-final/internal/config"
	"github.com/Onnywrite/lms-final/internal/domain"
	"log/slog"
)

func main() {
	const op = "main"

	cfg := config.MustLoadMain()

	logger := domain.MustSetupLoggerInDir(cfg.Env, cfg.LogsDir)

	log := logger.With(slog.String("op", op))

	application := app.NewMain(logger, cfg.Port, cfg.StaticDir)
	log.Debug("Starting the application")

	if err := application.Start(); err != nil {
		log.Error("Could not start the application")
		return
	}

	// graceful shutdown
	domain.ShutdownOnSignal(cfg.ShutdownTimeout, logger, application)
}
