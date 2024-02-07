package app

import (
	"fmt"
	"github.com/Onnywrite/lms-final/internal/services/calculator"
	"github.com/Onnywrite/lms-final/internal/services/restful"
	"log/slog"
)

type App struct {
	log      *slog.Logger
	server   *restful.Server
	restPort int
}

func New(
	logger *slog.Logger,
	dbConnect string,
	goroutinesCount int,
	port int) *App {
	// TODO: database (storage)
	//db := storage....
	calc := calculator.New(goroutinesCount /*, db*/)
	serv := restful.New(calc, port)

	return &App{
		log:      logger,
		server:   serv,
		restPort: port,
	}
}

func (a *App) MustStart() {
	if err := a.Start(); err != nil {
		panic(err)
	}
}

func (a *App) Start() error {
	const op = "app.Start"

	if err := a.server.Start(); err != nil {
		fmt.Errorf("%s: %s", op, err.Error())
		return err
	}
	a.log.Info("Server started")
	return nil
}

func (a *App) Stop() {
	a.server.Stop()
	a.log.Info("Server stopped")
}
