package restful

import "github.com/Onnywrite/lms-final/internal/services/calculator"

type Server struct {
	calc *calculator.Calculator
	port int
}

func New(calculator *calculator.Calculator, port int) *Server {
	return &Server{
		calc: calculator,
		port: port,
	}
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop() {

}
