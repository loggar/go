package slice1

import "fmt"

//lint:ignore U1000 // ignore unused variable warning
func arrayAppend() {
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
