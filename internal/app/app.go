package app

import (
	"log/slog"
)

type App struct {
	log      *slog.Logger
	dbToken  string
	goCount  int
	calcPort int
	restPort int
}

func New(
	logger *slog.Logger,
	dbToken string,
	goroutinesCount int,
	calculatorPort int,
	restfulPort int) *App {
	return &App{
		log:      logger,
		dbToken:  dbToken,
		goCount:  goroutinesCount,
		calcPort: calculatorPort,
		restPort: restfulPort,
	}
}

func (a *App) Start() error {
	const op = "app.Start"

	a.log.Debug("App started", slog.Int("calcPort", a.calcPort), slog.Int("restPort", a.restPort))
	return nil
}
