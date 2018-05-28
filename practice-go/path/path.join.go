package main

import (
	"fmt"
	"path"
)

func main() {
	base := "/home/bob"
	fmt.Println(path.Join(base, "work/go", "src/github.com"))
}
