package main

import (
	"fmt"
	"sync"
)

func main_2() {
	buffchan := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 1; i <= 2; i++ {
		go func(n int) {
			buffchan <- n
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(buffchan)

	for c := range buffchan {
		fmt.Println(c)
	}
}
