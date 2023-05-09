package main

import (
	"fmt"
	"strings"
)

func for_each() {
	// slice
	// var bookings = []string{}
	bookings := []string{}

	bookings = append(bookings, "Jane")
	bookings = append(bookings, "John")
	bookings = append(bookings, "Peter", "Tom")

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}

	fmt.Printf("fistNames: %v\n", firstNames)

}
