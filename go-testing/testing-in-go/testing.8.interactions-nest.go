package basics

// Train ...
type Train struct {
	engine *Engine
}

// NewTrain ...
func NewTrain() *Train {
	return &Train{
		engine: NewEngine(),
	}
}

// Go ...
func (t *Train) Go() {
	t.engine.Accel()
	t.engine.Accel()
}

// Stop ...
func (t *Train) Stop() {
	t.engine.Decel()
	t.engine.Decel()
}

// Speed ...
func (t *Train) Speed() int {
	return t.engine.Speed()
}
