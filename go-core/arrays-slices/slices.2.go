package main

import "fmt"

func main() {
	// slice
	// var bookings = []string{}
	bookings := []string{}

	bookings = append(bookings, "Jane")
	bookings = append(bookings, "John")
	bookings = append(bookings, "Peter", "Tom")

	fmt.Printf("The whole slice bookings: %v\n", bookings)
	fmt.Printf("slice bookings type: %T\n", bookings)
	fmt.Printf("slice bookings length: %v\n", len(bookings))

}
