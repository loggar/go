package basics

// Speeder ...
type Speeder interface {
	Speed() int
}

// Engine ...
type Engine interface {
	Speeder
	Accel()
	Decel()
}

type engine struct {
	speed int
}

// NewEngine ...
func NewEngine() Engine {
	return &engine{}
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	e.speed += 10
}

func (e *engine) Decel() {
	e.speed -= 10
}
