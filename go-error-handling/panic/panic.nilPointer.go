package main

import (
	"fmt"
)

// Shark ...
type Shark struct {
	Name string
}

// SayHello ...
func (s *Shark) SayHello() {
	fmt.Println("Hi! My name is", s.Name)
}

func panicNilPointer() {
	s := &Shark{"Sammy"}
	s = nil
	s.SayHello()
}
