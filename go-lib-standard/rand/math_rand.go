package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// randomNumber()
	randomNumberInRange()
}

func newSource() {
	source := rand.NewSource(time.Now().UnixNano())
	rand_source := rand.New(source)
	for i := 0; i < 5; i++ {
		rand_num := rand_source.Int()
		fmt.Println(rand_num)
	}
}

func randomNumber() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	num := r.Int()
	fmt.Printf("%d\n", num)
}

func randomNumberInRange() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		// generate an integer between 0 and 9
		dice_throw := r.Intn(10)
		// Move the Offset of 0
		fmt.Println(dice_throw + 1)
	}
}
