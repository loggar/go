package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	err := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", err)
}

func boom() error {
	return errors.New("barnacles")
}

func handlingErrors() {
	err := boom()

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("Anchors away!")
}
