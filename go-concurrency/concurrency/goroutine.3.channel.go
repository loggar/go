package main

import (
	"fmt"
	"strconv"
	"time"
)

func main_3() {
	c := make(chan string)
	go count_3("hello", c)

	for {
		msg, open := <-c
		if !open { // exit if channel is closed, or dead lock
			break
		}
		fmt.Println(msg)
	}
}

func count_3(str string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- strconv.Itoa(i) + string(' ') + str
		time.Sleep(time.Millisecond * 600)
	}

	close(c) // should close channel
}
