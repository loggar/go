package main

import "fmt"

func main() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for range sharks {
		sharks = append(sharks, "shark")
	}

	fmt.Printf("%q\n", sharks)
}
