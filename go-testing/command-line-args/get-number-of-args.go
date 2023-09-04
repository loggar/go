package main

import (
	"fmt"
	"os"
)

func main() {
	total_args := len(os.Args[1:])
	fmt.Println("Total Args:", total_args)
}
