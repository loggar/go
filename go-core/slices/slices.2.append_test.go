package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestSliceAppend(t *testing.T) {
	var percentages = []float64{78.8, 85.7, 94.4, 79.8}
	percentages = append(percentages, 60.5, 75.6)

	assert.Equal(t, 6, len(percentages), "len should equal %d", 6)
	assert.Equal(t, 8, cap(percentages), "cap should equal %d", 8)
}
