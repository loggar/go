package main

import "fmt"

type Creature_2 struct {
	Name string
	Type string

	password string
}

func main_2() {
	c := Creature_2{
		Name: "Sammy",
		Type: "Shark",

		password: "secret",
	}
	fmt.Println(c.Name, "the", c.Type)
	fmt.Println("Password is", c.password)
}
