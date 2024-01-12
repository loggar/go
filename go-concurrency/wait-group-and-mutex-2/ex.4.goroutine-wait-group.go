package main

import (
	"fmt"
	"sync"
	"time"
)

var globalValue4 int

func action4(i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	globalValue4 += i
	mutex.Unlock()

	time.Sleep(50 * time.Millisecond)

	wg.Done()
}

func main() {
	startTime := time.Now()

	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		fmt.Printf("Offset %d, Call Action\n", i)
		go action4(i, &wg, &mutex)
	}

	wg.Wait()

	delta := time.Now().Sub(startTime)
	fmt.Printf("Result is %d, done in %.3fs.\n", globalValue4, delta.Seconds())
}
