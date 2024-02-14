package orch

import (
	"context"
	"fmt"
	"github.com/Onnywrite/lms-final/internal/services/orch/handlers"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	log *slog.Logger
	srv *http.Server
}

func New(logger *slog.Logger, port int /*,interfaces...*/) *Server {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	mux := gin.New()
	mux.Use(gin.Recovery(), handlers.LogMiddleware(logger))

	mux.POST("/new", handlers.PostNew(nil))

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
			s.log.Error("", slog.String("op", op), slog.String("err", err.Error()))
		}
	}()

	s.log.Info("orch.Server started listening and serving new expressions")
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	const op = "orch.Stop"

	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Error("", slog.String("op", op), slog.String("err", err.Error()))
	}

	s.log.Info("restful.Server stopped its work")
}
