package main

import (
	"fmt"
	"time"
)

func main_1() {
	go count_1("hello") // sub-thread
	count_1("world")    // main thread, program is terminated when the main thread work finishes
}

func count_1(str string) {
	for i := 1; true; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 600)
	}
}
