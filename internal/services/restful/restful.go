package restful

import (
	"context"
	"fmt"
	"github.com/Onnywrite/lms-final/internal/domain/models"
	"github.com/gin-gonic/gin/binding"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	mux := gin.Default()
	binding.EnableDecoderDisallowUnknownFields = true

	mux.Static("/static", staticPath)
	mux.LoadHTMLFiles("./resources/index.html")

	mux.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	mux.POST("/new/", postNew)
	mux.GET("/status/", getStatus)
	// TODO: mux.HandleFunc("/powers/", handlePowers)
	// DEBUG-ONLY
	mux.GET("/ban/", func(c *gin.Context) {
		c.Status(http.StatusOK)
		os.Exit(0)
	})
	//

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
			fmt.Errorf("%s: %s", op, err.Error())
		}
	}()

	s.log.Info("restful.Server started listening and serving")
	return nil
}

func (s *Server) Stop(ctx context.Context) {
	const op = "restful.Stop"

	if err := s.srv.Shutdown(ctx); err != nil {
		fmt.Errorf("%s: %s", op, err.Error())
	}

	s.log.Info("restful.Server stopped its work")
}

func postNew(c *gin.Context) {
	body := models.Expression{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "could not parse JSON body request",
		})
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":     1567,
		"status": "calculating",
		"done":   "97.8%",
	})
}
