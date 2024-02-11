package calculator

import (
	"errors"
	"log/slog"

	"github.com/Onnywrite/lms-final/internal/domain/models"
)

type Calculator struct {
	log     *slog.Logger
	goCount int
}

var (
	InvalidTimingsErr = errors.New("timings cannot equal to zero")
)

func New(logger *slog.Logger, goroutinesCount int /*, storage.*/) *Calculator {
	// TODO: take DB instance as well
	logger.Debug("New calculator.Calculator was created")
	return &Calculator{
		log:     logger,
		goCount: goroutinesCount,
	}
}

func (c *Calculator) Start() error {
	c.log.Info("calculator.Calculator is started")
	return nil
}

func (c *Calculator) Stop() {
	c.log.Info("calculator.Calculator is stopped")
}

func (c *Calculator) ProcessExpression(expr *models.Expression) (
	processed models.ProcessedExpression, err error) {
	err = validateExpressionSettings(expr)

	return models.ProcessedExpression{}, nil
}

func validateExpressionSettings(e *models.Expression) error {
	if e.AdditionTime == 0 ||
		e.SubtractionTime == 0 ||
		e.MultiplicationTime == 0 ||
		e.DivisionTime == 0 {
		return InvalidTimingsErr
	}
	return nil
}

func parseExpression(e *models.Expression) ([]string, error) {
	return nil, nil
}
