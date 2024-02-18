package app

import (
	"context"
	"log/slog"

	"github.com/Onnywrite/lms-final/internal/services/restful"
)

type MainApp struct {
	log      *slog.Logger
	server   *restful.Server
	restPort int
}

func NewMain(
	logger *slog.Logger,
	port int,
	staticPath string) *MainApp {
	const op = "app.NewMain"
	log := logger.With(slog.String("op", op))

	serv := restful.New(logger, port, staticPath)

	log.Debug("MainApp was created")
	return &MainApp{
		log:      logger,
		server:   serv,
		restPort: port,
	}
}

func (a *MainApp) MustStart() {
	if err := a.Start(); err != nil {
		panic(err)
	}
}

func (a *MainApp) Start() error {
	const op = "MainApp.Start"
	log := a.log.With(slog.String("op", op))

	if err := a.server.Start(); err != nil {
		log.Error("", slog.String("err", err.Error()))
		return err
	}
	log.Info("MainApp started")
	return nil
}

func (a *MainApp) Stop(ctx context.Context, done chan<- error) {
	const op = "MainApp.Stop"
	log := a.log.With(slog.String("op", op))

	a.server.Stop(ctx, done)
	log.Info("MainApp stopped")
}
