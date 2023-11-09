package main

import (
	"fmt"
	"sync"
)

// receive-only channel
func receiver(ch <-chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("Received:", i)
	}
	wg.Done()
}

// send-only channel
func sender(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Sent:", i)
		ch <- i
	}
	close(ch)
	wg.Done()
}

func main_6() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go receiver(ch, &wg)
	go sender(ch, &wg)
	wg.Wait()
}
