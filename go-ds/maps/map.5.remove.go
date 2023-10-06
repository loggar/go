package main

import (
	"fmt"
)

func mapRemove() {
	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}
	delete(permissions, 16)
	fmt.Println(permissions)
}

func mapRemoveAll() {
	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}
	permissions = map[int]string{}
	fmt.Println(permissions)
}
