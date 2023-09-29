package main

import (
	"fmt"
)

func main_3() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println("Received:", num)
	}
}
