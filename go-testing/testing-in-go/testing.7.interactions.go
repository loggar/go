package basics

// Engine ...
type Engine struct {
	speed int
}

// NewEngine ...
func NewEngine() *Engine {
	return &Engine{}
}

// Speed ...
func (e *Engine) Speed() int {
	return e.speed
}

// Accel ...
func (e *Engine) Accel() {
	e.speed += 10
}

// Decel ...
func (e *Engine) Decel() {
	e.speed -= 10
}
