package main

import "fmt"

func slices() {
	mySlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mySlice)
	// [2 3 5 7 11 13]
	fmt.Println(mySlice[1:4])
	// [3 5 7]
	// missing low index implies 0
	fmt.Println(mySlice[:3])
	// [2 3 5]
	// missing high index implies len(s)
	fmt.Println(mySlice[4:])
	// [11 13]
}

func slicesWithMake() {
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	fmt.Printf("%q", cities)
	// ["Santa Monica" "Venice" "Los Angeles"]
}

func length() {
	cities := []string{
		"Santa Monica",
		"San Diego",
		"San Francisco",
	}
	fmt.Println(len(cities))
	// 3
	countries := make([]string, 42)
	fmt.Println(len(countries))
	// 42
}

func nilSlice() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("nil!")
	}
	// nil!
}
