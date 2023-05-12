package main

import "fmt"

type Book struct {
	pages int
	genre string
	title string
}

func main() {
	new_book := Book{120, "fiction", "Harry Potter"}
	fmt.Println(new_book)
	fmt.Printf("Type of new_book -> %T\n", new_book)
	book_ptr := &new_book
	book_ptr.title = "Games of Thrones"
	fmt.Println(new_book)
	fmt.Printf("Type of book_ptr -> %T\n", book_ptr)
}
