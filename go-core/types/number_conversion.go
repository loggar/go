package main

import (
	"fmt"
	"testing"
)

func numberConversion(t *testing.T) {
	var i = 42
	var f = float64(i)
	var u = uint(f)

	const fm = "%T(%v)\n"
	fmt.Printf(fm, i, i)
	fmt.Printf(fm, f, f)
	fmt.Printf(fm, u, u)
}
