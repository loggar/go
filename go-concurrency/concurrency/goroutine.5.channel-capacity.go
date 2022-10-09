package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)

	c <- "aaa"
	c <- "bbb"
	c <- "ccc" // channel is out of capacity

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
