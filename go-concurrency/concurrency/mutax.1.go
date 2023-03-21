package main

import (
	"fmt"
	"sync"
	"time"
)

var count_len = 70
var res []string
var mutex sync.Mutex

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < (count_len); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			dateStr := now.AddDate(0, 0, i).Format("20060102")
			mutex.Lock()
			res = append(res, dateStr)
			mutex.Unlock()
		}(&wg, i)
	}

	for i := 0; i < (count_len); i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			dateStr := now.AddDate(0, 0, i).Format("20060102")
			mutex.Lock()
			res = append(res, dateStr)
			mutex.Unlock()
		}(&wg, i)
	}

	wg.Wait()

	fmt.Println(fmt.Sprintf("should be %v ", (count_len*2)), len(res))
}
