package main

import "fmt"

func main() {
	q := new(Queue[string])
	q.enqueue("item-1")
	q.enqueue("item-2")
	q.enqueue("item-3")
	fmt.Println(q)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q)

	// q.enqueue(1) // cannot use 1 (untyped int constant) as string value in argument to q.enqueue
}

type Queue[T any] []T

func (q *Queue[T]) enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) dequeue() (T, bool) {
	if len(*q) == 0 {
		var v T
		return v, false
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, true
}
