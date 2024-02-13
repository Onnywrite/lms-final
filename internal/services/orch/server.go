package orch

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	log *slog.Logger
	srv *http.Server
}

// git add -p
type ExpressionProvider interface {
	Expression() error
}

func New(logger *slog.Logger, port int /*,interfaces...*/) *Server {
	mux := gin.Default()
	// TODO: Endpoints here

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	logger.Debug("New orch.Server is ready to handle")

	return &Server{
		log: logger,
		srv: srv,
	}
}

func (s *Server) Start() error {
	const op = "orch.Start"

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Errorf("%s: %s", op, err.Error())
		}
	}()

	s.log.Info("orch.Server started listening and serving new expressions")
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	const op = "orch.Stop"

	if err := s.srv.Shutdown(ctx); err != nil {
		fmt.Errorf("%s: %s", op, err.Error())
	}

	s.log.Info("restful.Server stopped its work")
}
