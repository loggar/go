package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum=%d\n", sum)

	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}
