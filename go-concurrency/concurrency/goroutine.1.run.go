package main

import (
	"fmt"
	"time"
)

func main() {
	go count("hello") // sub-thread
	count("world")    // main thread, program is terminated when the main thread work finishes
}

func count(str string) {
	for i := 1; true; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 600)
	}
}
