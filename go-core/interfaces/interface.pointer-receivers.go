package main

import "fmt"

type Submersible interface {
	Dive()
}

type Shark struct {
	Name string

	isUnderwater bool
}

func (s Shark) String() string {
	if s.isUnderwater {
		return fmt.Sprintf("%s is underwater", s.Name)
	}
	return fmt.Sprintf("%s is on the surface", s.Name)
}

func (s *Shark) Dive() {
	s.isUnderwater = true
}

func submerge(s Submersible) {
	s.Dive()
}

func main() {
	s := &Shark{
		Name: "Sammy",
	}

	fmt.Println(s)

	submerge(s)

	fmt.Println(s)
}
