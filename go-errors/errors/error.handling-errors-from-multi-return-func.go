package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize2(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func fromMultiReturnFn() {
	_, err := capitalize2("")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}
	fmt.Println("Success!")
}
