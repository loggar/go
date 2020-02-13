package main

import "fmt"

func main() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i, shark := range sharks {
		fmt.Println(i, shark)
	}
}
