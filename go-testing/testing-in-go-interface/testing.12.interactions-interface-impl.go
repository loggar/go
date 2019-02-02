package basics

// FakeEngine ...
type FakeEngine struct {
	AccelCalls int
	DecelCalls int
}

// NewFakeEngine ...
func NewFakeEngine() *FakeEngine {
	return &FakeEngine{}
}

// Accel ...
func (e *FakeEngine) Accel() {
	e.AccelCalls++
}

// Decel ...
func (e *FakeEngine) Decel() {
	e.DecelCalls++
}

// Speed ...
func (e *FakeEngine) Speed() int {
	return 0
}

// Train ...
type Train struct {
	engine Engine
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
