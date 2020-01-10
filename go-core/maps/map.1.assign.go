package main

import "fmt"

func main() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)

	fmt.Println(sammy["animal"])
	// returns shark

	fmt.Println(sammy["color"])
	// returns blue

	fmt.Println(sammy["location"])
	// returns ocean
}
