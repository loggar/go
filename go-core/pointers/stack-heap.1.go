package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() *person {
	p := person{name: "a", age: 50}
	fmt.Printf("initPerson: %p\n", &p)
	return &p
}

func main_a_1() {
	m := initPerson()
	fmt.Printf("main: %p\n", m)
}
