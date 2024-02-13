package app

import (
	"fmt"
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
	serv := restful.New(logger, port, staticPath)

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
	const op = "app.Start"

	if err := a.server.Start(); err != nil {
		fmt.Errorf("%s: %s", op, err.Error())
		return err
	}
	a.log.Info("Server started")
	return nil
}

func (a *MainApp) Stop() {
	a.server.Stop()
	a.log.Info("Server stopped")
}
