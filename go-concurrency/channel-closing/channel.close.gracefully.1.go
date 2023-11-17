package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// https://go101.org/article/channel-closing.html
// 1. M receivers, one sender, the sender says "no more sends" by closing the data channel

func main_1() {
	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20
	log.SetFlags(0)

	// ...
	const Max = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// the sender
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// The only sender can close the
				// channel at any time safely.
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// Receive values until dataCh is
			// closed and the value buffer queue
			// of dataCh becomes empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
