package restful

import (
	"fmt"
	"github.com/Onnywrite/lms-final/internal/domain/models"
	"github.com/gin-gonic/gin/binding"
	"log/slog"
	"net/http"
	"os"

	"github.com/Onnywrite/lms-final/internal/services/calculator"
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
	log  *slog.Logger
	calc *calculator.Calculator
	mux  *gin.Engine
	port int
}

func New(logger *slog.Logger, calculator *calculator.Calculator, port int, staticPath string) *Server {
	mux := gin.Default()
	binding.EnableDecoderDisallowUnknownFields = true

	mux.Static("/static", staticPath)
	mux.Use(func(c *gin.Context) {
		c.Set("calc", calculator)
	})
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

func postNew(c *gin.Context) {
	body := models.Expression{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "could not parse JSON body request",
		})
		return
	}

	if calcAny, exists := c.Get("calc"); exists {
		expr, err := calcAny.(*calculator.Calculator).ProcessExpression(&body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "could not process your expression",
			})
			return
		}

		c.JSON(http.StatusAccepted, &expr)
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
