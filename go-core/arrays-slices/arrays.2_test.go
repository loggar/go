package array1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray1(t *testing.T) {
	var bookings = [50]string{}

	bookings[0] = "Jane"
	bookings[1] = "John"
	bookings[2] = "Peter"
	bookings[3] = "Tom"

	fmt.Printf("The whole array bookings: %v\n", bookings)
	fmt.Printf("array bookings type: %T\n", bookings)
	fmt.Printf("array bookings length: %v\n", len(bookings))

	length := len(bookings)
	expected := 50
	assert.Equal(t, expected, length, "len(%v) should equal %f", bookings, expected)
}
