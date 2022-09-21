package main

import "fmt"

func main() {
	addAnonymous := func(nums ...int) (result int) {
		fmt.Println("anonymous function")
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}

	addAnonymous(1, 2, 3, 4, 5)
}

func addAnonymous(nums ...int) (result int) {
	fmt.Println("defined function")
	return nums[0]
}
