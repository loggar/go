package main

import (
	"fmt"
	"sync"
	"time"
)

func main_2() {
	var wg sync.WaitGroup
	wg.Add(1) // I have 1 goroutine to wait for.

	go func() {
		count_2("hello")
		wg.Done()
	}()

	wg.Wait() // this will wait until the wait group counts 0.
}

func count_2(str string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 600)
	}
}
