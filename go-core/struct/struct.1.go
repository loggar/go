package main

import "fmt"

type Creature struct {
	Name string
	Type string
}

func main() {
	c := Creature{"Sammy", "Shark"}
	fmt.Println(c.Name, "the", c.Type)
}
