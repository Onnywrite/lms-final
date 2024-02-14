package restful

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Server has these endpoints:
//
// GET / with the index page;
//
// POST new/ must be with body. Adding new expression for calculating it.
// could send 200 OK, 400 Invalid expression, 500 Internal error. If OK, it'll send the id of
// your expression;
//
// GET status/ returns list of statuses of all expressions;
//
// GET status/(id) with status of expression with the id;
//
// ....
type Server struct {
	log *slog.Logger
	srv *http.Server
}

func New(logger *slog.Logger, port int, staticPath string) *Server {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	mux := gin.Default()
	binding.EnableDecoderDisallowUnknownFields = true

	mux.Static("/static", staticPath)
	mux.LoadHTMLFiles("./resources/index.html")

	mux.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	mux.GET("/status", getStatus)
	mux.GET("/servers", getServers)
	mux.POST("/new", postNew)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	logger.Debug("New restful.Server is ready to handle")

	return &Server{
		log: logger,
		srv: srv,
	}
}

func (s *Server) Start() error {
	const op = "restful.Start"

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Error("", slog.String("op", op), slog.String("err", err.Error()))
		}
	}()

	s.log.Info("restful.Server started listening and serving")
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	const op = "restful.Stop"

	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Error("", slog.String("op", op), slog.String("err", err.Error()))
	}

	s.log.Info("restful.Server stopped its work")
}
