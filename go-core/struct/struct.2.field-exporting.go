package main

import "fmt"

type Creature struct {
	Name string
	Type string

	password string
}

func main() {
	c := Creature{
		Name: "Sammy",
		Type: "Shark",

		password: "secret",
	}
	fmt.Println(c.Name, "the", c.Type)
	fmt.Println("Password is", c.password)
}
