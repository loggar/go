package main

import (
	"fmt"
	"time"
)

var globalValue1 int

func action1(i int) {
	globalValue1 += i
	time.Sleep(50 * time.Millisecond)
}

func main_1() {
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		fmt.Printf("Offset %d, Call Action\n", i)
		action1(i)
	}

	delta := time.Now().Sub(startTime)
	fmt.Printf("Result is %d, done in %.3fs.\n", globalValue1, delta.Seconds())
}
