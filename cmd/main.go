package main

import (
	"github.com/Onnywrite/lms-final/internal/app"
	"github.com/Onnywrite/lms-final/internal/config"
	"github.com/Onnywrite/lms-final/internal/domain"
)

func main() {
	cfg := config.MustLoadMain()

	log := domain.MustSetupLoggerInDir(cfg.Env, cfg.LogsDir)

	application := app.NewMain(log, cfg.Port, cfg.StaticDir)
	log.Debug("Starting the application")
	go application.MustStart()

	// graceful shutdown
	domain.ShutdownOnSignal(cfg.ShutdownTimeout, log, application)
}
