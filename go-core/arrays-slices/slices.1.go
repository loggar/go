package array1

import "fmt"

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
func slicesWithMake() {
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
