package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomShuffleLetters() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	letters := "abcdefghijklmnopqrstuvwxyz"
	shuffled := r.Perm(len(letters))

	result := make([]byte, len(letters))
	for i, randIndex := range shuffled {
		result[i] = letters[randIndex]
	}
	rand_str := string(result)
	fmt.Println(rand_str)
	// random string of length 10
	fmt.Println(rand_str[:10])
}
