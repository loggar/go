package main

import "fmt"

type Book struct {
	Id   int64
	Name string
}

func for_struct() {
	list := []*Book{
		{
			Id:   123,
			Name: "name123",
		},

		{
			Id:   124,
			Name: "name124",
		},
	}

	fmt.Printf("list (%T): %v\n", list, list)

	var book *Book
	for i, b := range list {
		fmt.Println(i, b)
		if b.Id == 124 {
			book = b
		}
	}

	fmt.Printf("book: %v", book)
}
