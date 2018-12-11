package basics

import "testing"

func expectSpeed(t *testing.T, e *Engine, s int) {
	actual := e.Speed()
	if actual != s {
		t.Fatal("expected", s, "got", actual)
	}
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	expectSpeed(t, e, 0)
	e.Accel()
	expectSpeed(t, e, 10)
	e.Decel()
	expectSpeed(t, e, 0)
}
