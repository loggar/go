package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//lint:ignore U1000 // ignore unused variable warning
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

//lint:ignore U1000 // ignore unused variable warning
func slicesFromMake() {
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	fmt.Printf("%q", cities)
	// ["Santa Monica" "Venice" "Los Angeles"]
}

//lint:ignore U1000 // ignore unused variable warning
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

//lint:ignore U1000 // ignore unused variable warning
func nilSlice() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("nil!")
	}
	// nil!
}

//lint:ignore U1000 // ignore unused variable warning
func append1() {
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)
}

//lint:ignore U1000 // ignore unused variable warning
func append2() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}

func TestSliceMakeLenCap(t *testing.T) {
	var langs = make([]string, 3, 5)

	langs[0], langs[1], langs[2] = "Python", "Go", "Javascript"

	assert.Equal(t, 3, len(langs), "len should equal %d", 3)
	assert.Equal(t, 5, cap(langs), "cap should equal %d", 5)

	langs = append(langs, "Java", "Kotlin", "PHP")

	assert.Equal(t, 6, len(langs), "len should equal %d", 6)
	assert.Equal(t, 10, cap(langs), "cap should equal %d", 10)
}
