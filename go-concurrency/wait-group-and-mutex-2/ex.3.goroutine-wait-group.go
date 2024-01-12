package main

import (
	"fmt"
	"sync"
	"time"
)

var globalValue3 int

func action3(i int, wg *sync.WaitGroup) {
	globalValue3 += i
	time.Sleep(50 * time.Millisecond)

	wg.Done()
}

func main_3() {
	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		fmt.Printf("Offset %d, Call Action\n", i)
		go action3(i, &wg)
	}

	wg.Wait()

	delta := time.Now().Sub(startTime)
	fmt.Printf("Result is %d, done in %.3fs.\n", globalValue3, delta.Seconds())
}
