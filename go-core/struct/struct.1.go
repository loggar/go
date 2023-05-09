package main

import "fmt"

type Creature_1 struct {
	Name string
	Type string
}

func main_1() {
	c := Creature_1{"Sammy", "Shark"}
	fmt.Println(c.Name, "the", c.Type)
}
