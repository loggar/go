package main

import (
	"fmt"
)

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type snake struct {
	Type      string
	Poisonous bool
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound: %v", c.Sound)
}

func testmain() {
	var a animal
	a = snake{Poisonous: true}
	fmt.Println(a.description())
	a = cat{Sound: "Meow!!!"}
	fmt.Println(a.description())
}

//=> Poisonous: true
//=> Sound: Meow!!!
