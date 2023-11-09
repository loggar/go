package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
	default:
		fmt.Println("Channel Blocking")
	}

	fmt.Println(<-ch)
}

// Channel Blocking
// 1
