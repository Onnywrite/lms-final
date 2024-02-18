package restful

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Onnywrite/lms-final/internal/domain/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Server struct {
	log *slog.Logger
	srv *http.Server
}

func New(logger *slog.Logger, port int, staticPath string) *Server {
	const op = "restful.New"
	log := logger.With(slog.String("op", op))

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	binding.EnableDecoderDisallowUnknownFields = true

	mux := gin.New()
	mux.Use(middleware.Recover(logger), middleware.Log(logger))

	mux.Static("/static", staticPath)
	mux.LoadHTMLFiles("./resources/index.html")

	mux.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Debug("ready to handle")

	return &Server{
		log: logger,
		srv: srv,
	}
}

func (s *Server) Start() error {
	const op = "restful.Start"
	log := s.log.With(slog.String("op", op))

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("", slog.String("op", op), slog.String("err", err.Error()))
		}
	}()

	log.Info("started listening and serving")
	return nil
}

func (s *Server) Stop(ctx context.Context, done chan<- error) {
	const op = "restful.Server.Stop"
	log := s.log.With(slog.String("op", op))

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Error("", slog.String("op", op), slog.String("err", err.Error()))
		done <- err
		return
	}

	log.Info("stopped its work")
}
