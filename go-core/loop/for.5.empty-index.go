package main

import "fmt"

func for_range_empty_index() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for _, shark := range sharks {
		fmt.Println(shark)
	}
}
