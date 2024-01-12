package main

import (
	"fmt"
	"time"
)

var globalValue2 int

func action2(i int) {
	globalValue2 += i
	time.Sleep(50 * time.Millisecond)
}

func main_2() {
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		fmt.Printf("Offset %d, Call Action\n", i)
		go action2(i)
	}

	delta := time.Now().Sub(startTime)
	fmt.Printf("Result is %d, done in %.3fs.\n", globalValue2, delta.Seconds())
}
