package main

import "fmt"

func test_5() {
	var creature *Creature
	creature = &Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}
