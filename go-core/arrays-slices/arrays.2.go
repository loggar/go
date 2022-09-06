package main

import "fmt"

func main() {
	// var bookings = [50]string{"Jane", "John", "Peter"}
	var bookings = [50]string{}

	bookings[0] = "Jane"
	bookings[1] = "John"
	bookings[2] = "Peter"
	bookings[3] = "Tom"

	fmt.Printf("The whole array bookings: %v\n", bookings)
	fmt.Printf("array bookings type: %T\n", bookings)
	fmt.Printf("array bookings length: %v\n", len(bookings))

}
