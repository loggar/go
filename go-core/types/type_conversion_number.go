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

func intConversion1() {
	var index int8 = 15

	var bigIndex int32

	bigIndex = int32(index)

	fmt.Println(bigIndex)
}

func intConversion2() {
	var big int64 = 64

	var little int8

	little = int8(big)

	fmt.Println(little)
}

func intToFloat() {
	var x int64 = 57

	var y = float64(x)

	fmt.Printf("%.2f\n", y)
}

func floatToInt1() {
	var f = 390.8
	var i = int(f)

	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)
}

func floatToInt2() {
	b := 125.0
	c := 390.8

	fmt.Println(int(b))
	fmt.Println(int(c))
}

func throughDivision() {
	a := 5 / 2
	fmt.Println(a)

	b := 5.0 / 2
	fmt.Println(b)
}
