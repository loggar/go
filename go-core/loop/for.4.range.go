package main

import "fmt"

func for_slice_range() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i, shark := range sharks {
		fmt.Println(i, shark)
	}
}
