package calculator

type Calculator struct {
	goCount int
}

func New(goroutinesCount int /*, storage.*/) *Calculator {
	// TODO: take DB instance as well
	return &Calculator{goCount: goroutinesCount}
}

func (c *Calculator) Start() {

}

func (c *Calculator) Stop() {

}
