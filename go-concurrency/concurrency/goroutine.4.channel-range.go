package main

import (
	"fmt"
	"strconv"
	"time"
)

func main_4() {
	c := make(chan string)
	go count_4("hello", c)

	for msg := range c { // syntactic sugar
		fmt.Println(msg)
	}
}

func count_4(str string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- strconv.Itoa(i) + string(' ') + str
		time.Sleep(time.Millisecond * 600)
	}

	close(c) // should close channel
}
