package main

import (
	"bytes"
	"testing"
)

func TestSave(t *testing.T) {
	p1 := &Page{Title: "testPage2", Body: []byte("body1")}
	err := p1.save()

	if err != nil {
		t.Errorf("fail, %s", err)
	}
}

func TestLoadPage(t *testing.T) {
	p1 := &Page{Title: "testPage2", Body: []byte("body1")}
	err := p1.save()

	if err != nil {
		t.Errorf("fail, %s", err)
	}

	p2, _ := loadPage("testPage2")

	if !bytes.Equal(p2.Body, p1.Body) {
		t.Errorf("fail, got: %s, want: %s.", p2.Body, p1.Body)
	}
}
