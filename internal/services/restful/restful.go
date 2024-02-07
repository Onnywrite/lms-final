package restful

import (
	"fmt"
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
	mux  *http.ServeMux
	port int
}

func New(logger *slog.Logger, calculator *calculator.Calculator, port int) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/new/", handleNew)
	mux.HandleFunc("/status/", handleStatus)
	mux.HandleFunc("/setting/", handleSetting)
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
	if err := s.calc.Start(); err != nil {
		return err
	}

	http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)

	s.log.Info("restful.Server started listening and serving")
	return nil
}

func (s *Server) Stop() {
	s.calc.Stop()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"id":"hello from Index"}`))
	w.WriteHeader(http.StatusOK)
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"id":"hello from New"}`))
	w.WriteHeader(http.StatusOK)
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"id":"hello from Status", "expression":"status2+2"}`))
	w.WriteHeader(http.StatusOK)
}

func handleSetting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"id":"hello from Setting"}`))
	w.WriteHeader(http.StatusOK)
}
