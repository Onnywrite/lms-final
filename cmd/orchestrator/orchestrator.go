package main

import (
	"github.com/Onnywrite/lms-final/internal/app"
	"github.com/Onnywrite/lms-final/internal/config"
	"github.com/Onnywrite/lms-final/internal/domain"
)

func main() {
	cfg := config.MustLoadOrchestrator()

	log := domain.MustSetupLoggerInDir(cfg.Env, cfg.LogsDir)

	application := app.NewOrchestrator(log, cfg.Port, cfg.DbConnection)
	log.Debug("Starting the orchestrator")
	go application.MustStart()

	// graceful shutdown
	domain.ShutdownOnSignal(cfg.ShutdownTimeout, log, application)
}
