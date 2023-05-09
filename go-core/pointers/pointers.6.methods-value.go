/*
Method Pointer Receivers
*/
package main

import "fmt"

func (c Creature) Reset() {
	c.Species = ""
}

func test_6() {
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	creature.Reset()
	fmt.Printf("2) %+v\n", creature)
}
