package orch

import (
	"log/slog"
)

type Server struct {
	log *slog.Logger
}

type ExpressionProvider interface {
	Expression() error
}

func New(logger *slog.Logger, port int /*, db *storage.Storage*/) *Server {

}
