package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // I have 1 goroutine to wait for.

	go func() {
		count("hello")
		wg.Done()
	}()

	wg.Wait() // this will wait until the wait group counts 0.
}

func count(str string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 600)
	}
}
