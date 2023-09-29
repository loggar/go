package main

import (
	"fmt"
)

func main_1() {
	ch := make(chan string)
	go func() {
		message := "Hello, Gophers!"
		ch <- message
	}()
	fmt.Println(<-ch) // Hello, Gophers!
	fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
}
