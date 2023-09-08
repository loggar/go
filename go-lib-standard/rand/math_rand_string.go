package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Random string
	randomLowerCase := make([]rune, 6)
	randomUpperCase := make([]rune, 6)
	for i := range randomLowerCase {
		randomLowerCase[i] = rune(r.Intn(26) + 97)
		randomUpperCase[i] = rune(r.Intn(26) + 65)
	}
	randomLowerCaseStr := string(randomLowerCase)
	randomUpperCaseStr := string(randomUpperCase)

	fmt.Println(randomLowerCase)
	fmt.Println(randomLowerCaseStr)
	fmt.Println(randomUpperCase)
	fmt.Println(randomUpperCaseStr)
}
