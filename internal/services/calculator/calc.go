package calculator

import "log/slog"

type Calculator struct {
	log     *slog.Logger
	goCount int
}

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
