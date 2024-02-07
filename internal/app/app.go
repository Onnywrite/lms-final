package app

import (
	"log/slog"
)

type App struct {
	log     *slog.Logger
	dbConn  string
	goCount int
	Port    int
}

func New(
	logger *slog.Logger,
	dbConnect string,
	goroutinesCount int,
	Port int) *App {
	return &App{
		log:     logger,
		dbConn:  dbConnect,
		goCount: goroutinesCount,
		Port:    Port,
	}
}

func (a *App) MustStart() {
	err := a.Start()
	if err != nil {
		panic(err)
	}
}

func (a *App) Start() error {
	const op = "app.MustStart"

	//a.log.Info("App started", slog.Int("calcPort", a.Port), slog.Int("restPort", a.Port))
	return nil
}

func (a *App) Stop() error {
	const op = "app.Start"

	//a.log.Debug("App started", slog.Int("calcPort", a.Port), slog.Int("restPort", a.Port))
	return nil
}
