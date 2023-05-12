package main

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main_b_1() {
	x := 3
	y := 6
	k := &x
	p := &y
	fmt.Printf("Before swapping : x = %d and y = %d.\n", x, y)
	swap(k, p)
	fmt.Printf("After swapping  : x = %d and y = %d.\n", x, y)
}
