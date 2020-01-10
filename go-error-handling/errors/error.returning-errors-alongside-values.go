package main

import (
	"errors"
	"fmt"
	"strings"
)

func capitalize1(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func returnValues() {
	name, err := capitalize1("sammy")
	if err != nil {
		fmt.Println("Could not capitalize:", err)
		return
	}

	fmt.Println("Capitalized name:", name)
}
