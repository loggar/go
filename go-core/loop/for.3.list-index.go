package main

import "fmt"

func for_slice() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	}
}
