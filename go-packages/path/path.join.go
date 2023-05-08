package main

import (
	"fmt"
	"path"
)

func pathJoin() {
	base := "/home/bob"
	fmt.Println(path.Join(base, "work/go", "src/github.com"))
}
