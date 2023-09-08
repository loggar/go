package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomShuffleArray() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println(r.Perm(10))

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before shuffle:", arr)
	r.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("After shuffle:", arr)
}
