package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomFloat() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rand_float32 := r.Float32()
	fmt.Println(rand_float32)

	rand_float64 := r.Float64()
	fmt.Println(rand_float64)

	rand_exp_float := r.ExpFloat64()
	fmt.Println(rand_exp_float)

	rand_norm_float := r.NormFloat64()
	fmt.Println(rand_norm_float)

	for i := 0; i < 5; i++ {
		rand_float := r.Float32()
		fmt.Println(rand_float)
	}
}
