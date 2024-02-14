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
	dbConnection string) *OrchestratorApp {
	// db := mongo.New(...)

	serv := orch.New(logger, port /*,db,db,...*/)

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
	const op = "app.Start"

	if err := a.server.Start(); err != nil {
		a.log.Error("", slog.String("op", op), slog.String("err", err.Error()))
		return err
	}
	a.log.Info("Server started")
	return nil
}

func (a *OrchestratorApp) Stop(ctx context.Context) {
	a.server.Stop(ctx)
	a.log.Info("Server stopped")
}
