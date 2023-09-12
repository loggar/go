package main

import (
	"fmt"
)

type Stack2[T any] struct {
	items []T
}

func NewStack2[T any]() *Stack2[T] {
	return &Stack2[T]{}
}

func (s *Stack2[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack2[T]) Pop() T {
	if len(s.items) == 0 {
		panic("Stack2 is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func usage_struct_stack() {
	intStack := NewStack2[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println("Integer Stack")
	fmt.Println(intStack)
	intStack.Pop()
	intStack.Pop()
	fmt.Println(intStack)

	// without the NewStack2 method
	strStack := Stack2[string]{}
	strStack.Push("c")
	strStack.Push("python")
	strStack.Push("mojo")
	fmt.Println("String Stack:")
	fmt.Println(strStack)
	strStack.Pop()
	fmt.Println(strStack)
}
