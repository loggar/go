package main

import "fmt"

// Artist2 struct
type Artist2 struct {
	Name, Genre string
	Songs       int
}

func newRelease2(a *Artist2) int {
	a.Songs++
	return a.Songs
}
func passedByReference() {
	me := &Artist2{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease2(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
