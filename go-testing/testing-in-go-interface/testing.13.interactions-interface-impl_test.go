package basics

import "testing"

func trainWithFakeEngine() (*Train, *FakeEngine) {
	t := NewTrain()
	e := NewFakeEngine()
	t.engine = e
	return t, e
}

func TestTrainGo(t *testing.T) {
	tr, e := trainWithFakeEngine()
	tr.Go()
	if e.AccelCalls != 2 {
		t.Error("expected", 2, "got", e.AccelCalls)
	}
}

func TestTrainStop(t *testing.T) {
	tr, e := trainWithFakeEngine()
	tr.Stop()
	if e.DecelCalls != 2 {
		t.Error("expected", 2, "got", e.DecelCalls)
	}
}
