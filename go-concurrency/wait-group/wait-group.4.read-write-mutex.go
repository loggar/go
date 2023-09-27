package main

import (
	"fmt"
	"sync"
	"time"
)

func reader4(id int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	if !mutex.TryLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	defer mutex.Unlock()
	fmt.Printf("Reader %d: read count %d\n", id, *count)
}

func writer4(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf("Writer %d: wrote count %d\n", id, *count)
}

func main_4() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.Mutex
	readers := 5
	writers := 3
	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go reader4(i, &count, &wg, &mutex)
	}

	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go writer4(i, 1, &count, &wg, &mutex)
	}
	wg.Wait()

}
