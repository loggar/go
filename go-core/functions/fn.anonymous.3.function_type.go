package main

import "fmt"

func main() {
	is_divisible := func(x int, y int) bool {
		var res bool
		if x%y == 0 {
			res = true
		} else {
			res = false
		}
		fmt.Println(res)
		return res
	}

	fmt.Printf("%T\n", is_divisible)
	is_divisible(10, 5)
	is_divisible(33, 5)

	divisibility_check := is_divisible(45, 5)
	fmt.Printf("%T : %t\n", divisibility_check, divisibility_check)
}
