package restful

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"

	"github.com/Onnywrite/lms-final/internal/services/calculator"
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
// GET/PUT/POST setting/(id) gets or edits calculation settings of the
// expression with given id;
//
// ....
type Server struct {
	log  *slog.Logger
	calc *calculator.Calculator
	mux  *gin.Engine
	port int
}

func New(logger *slog.Logger, calculator *calculator.Calculator, port int) *Server {
	mux := gin.Default()

	mux.Static("/static", "./resources")

	mux.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	mux.POST("/new/", getNew)
	mux.GET("/status/", getStatus)
	mux.GET("/setting/", getSetting)
	mux.POST("/setting/", postSetting)
	mux.PUT("/setting/", putSetting)
	// TODO: mux.HandleFunc("/powers/", handlePowers)
	logger.Debug("New restful.Server is ready to handle")

	return &Server{
		log:  logger,
		calc: calculator,
		port: port,
		mux:  mux,
	}
}

func (s *Server) Start() error {
	const op = "restful.Start"

	if err := s.calc.Start(); err != nil {
		return err
	}

	go func() {
		if err := s.mux.Run(fmt.Sprintf(":%d", s.port)); err != nil {
			fmt.Errorf("%s: %s", op, err.Error())
		}
	}()

	s.log.Info("restful.Server started listening and serving")
	return nil
}

func (s *Server) Stop() {
	s.calc.Stop()
}

func getNew(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": 1567,
	})
}

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":     1567,
		"status": "calculating",
		"done":   "97.8%",
	})
}

func getSetting(c *gin.Context) {

}

func postSetting(c *gin.Context) {

}

func putSetting(c *gin.Context) {

}
