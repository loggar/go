package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer_2 interface {
	String() string
}

func test_stringer_2() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print(b)
}

func Print(s Stringer_2) {
	fmt.Println(s.String())
}
