package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func goroutine() {
	time.Sleep(5 * time.Second)
	fmt.Printf("done")
	wg.Done()
}

func main() {
	wg.Add(1)

	go goroutine()

	wg.Wait()
}
