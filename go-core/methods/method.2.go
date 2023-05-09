package main

import "fmt"

func (c Creature) Greet2() Creature {
	fmt.Printf("%s says %s!\n", c.Name, c.Greeting)
	return c
}

func (c Creature) SayGoodbye(name string) {
	fmt.Println("Farewell", name, "!")
}

func test_2() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet2().SayGoodbye("gophers")

	Creature.SayGoodbye(Creature.Greet2(sammy), "gophers")
}
