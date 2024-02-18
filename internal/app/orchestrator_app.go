package app

import (
	"context"
	"log/slog"

	"github.com/Onnywrite/lms-final/internal/services/orch"
)

type OrchestratorApp struct {
	log    *slog.Logger
	server *orch.Server
	port   int
}

func NewOrchestrator(
	logger *slog.Logger,
	port int,
	// TODO: storage
	dbConnection string,
	allowOrigin []string) *OrchestratorApp {
	const op = "app.NewOrchestrator"
	log := logger.With(slog.String("op", op))

	// db := mongo.New(...)

	serv := orch.New(logger, port, allowOrigin /*,db,db,...*/)

	log.Debug("was created")

	return &OrchestratorApp{
		log:    logger,
		server: serv,
		port:   port,
	}
}

func (a *OrchestratorApp) MustStart() {
	if err := a.Start(); err != nil {
		panic(err)
	}
}

func (a *OrchestratorApp) Start() error {
	const op = "app.OrchestratorApp.Start"
	log := a.log.With(slog.String("op", op))

	if err := a.server.Start(); err != nil {
		log.Error("", slog.String("err", err.Error()))
		return err
	}
	log.Info("started")
	return nil
}

func (a *OrchestratorApp) Stop(ctx context.Context, done chan<- error) {
	const op = "app.OrchestratorApp.Stop"
	log := a.log.With(slog.String("op", op))

	a.server.Stop(ctx, done)
	log.Info("stopped")
}
