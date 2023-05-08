package main

import "fmt"

func main_worker_1() {
	max := 45
	jobs := make(chan int, max)
	results := make(chan int, max)

	go worker(jobs, results, 1)
	go worker(jobs, results, 2)
	go worker(jobs, results, 3)
	go worker(jobs, results, 4)
	go worker(jobs, results, 5)

	for i := 0; i < max; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < max; i++ {
		fmt.Printf("main result received %v\n", <-results)
	}
}

func worker(jobs <-chan int, results chan<- int, workerId int) {
	for n := range jobs {
		r := fib(n)
		fmt.Printf("worker %v fib(%v), result %v\n", workerId, n, r)
		results <- r
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
