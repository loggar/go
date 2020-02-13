package main

import "fmt"

func main() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for _, shark := range sharks {
		fmt.Println(shark)
	}
}
