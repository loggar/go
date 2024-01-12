package main

import (
	"fmt"
	"sync"
	"time"
)

func reader5(id int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.RUnlock()

	if !mutex.TryRLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	fmt.Printf("Reader %d: read count %d\n", id, *count)

}

func writer5(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()

	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf("Writer %d: wrote count %d\n", id, *count)
}

// go run --race rwmutex.go
func main_5() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	readers := 5
	writers := 3
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go reader5(i, &count, &wg, &mutex)
	}

	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go writer5(i, 1, &count, &wg, &mutex)
	}
	wg.Wait()

}
