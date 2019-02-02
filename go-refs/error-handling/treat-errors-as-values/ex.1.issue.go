package main

import "fmt"

func write(s string) int {
	fmt.Printf("%s\n", s)
	return len(s)
}

func main() {
	var count, n int
	n = write("sample text 1")
	count += n

	// error checking
	if count >= 32 {
		fmt.Printf("Written length : %d\n", count)
		return
	}

	n = write("sample text 2")
	count += n

	if count >= 32 {
		fmt.Printf("Written length : %d\n", count)
		return
	}

	n = write("sample text 3")
	count += n

	if count >= 32 {
		fmt.Printf("Written length : %d\n", count)
		return
	}

	n = write("sample text 4")
	count += n

	if count >= 32 {
		fmt.Printf("Written length : %d\n", count)
		return
	}

}
