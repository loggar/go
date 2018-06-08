package main

import "fmt"

func maps() {
	celebs := map[string]int{
		"Nicolas Cage":       50,
		"Selena Gomez":       21,
		"Jude Law":           41,
		"Scarlett Johansson": 29,
	}
	fmt.Printf("%#v", celebs)
}

// Vertex ...
type Vertex struct {
	Lat, Long float64
}

func mapsWithMake() {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		// same as "Bell Labs": Vertex{40.68433, -74.39967}
		"Google": {37.42202, -122.08408},
	}

	fmt.Println(m)
}

/* mutating maps
m[key] = elem
elem = m[key]
delete(m, key)
elem, ok = m[key]
*/
