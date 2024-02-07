package calculator

import "log/slog"

type Calculator struct {
	log     *slog.Logger
	goCount int
}

func New(logger *slog.Logger, goroutinesCount int /*, storage.*/) *Calculator {
	// TODO: take DB instance as well
	return &Calculator{
		log:     logger,
		goCount: goroutinesCount,
	}
}

func (c *Calculator) Start() error {
	return nil
}

func (c *Calculator) Stop() {

}
