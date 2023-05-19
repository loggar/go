package main

import "fmt"

func bitwise() {
	x := 3
	y := 5
	// 3 -> 011
	// 5 -> 101
	fmt.Println("X AND Y = ", x&y)
	fmt.Println("X OR Y = ", x|y)
	fmt.Println("X EXOR Y = ", x^y)
	fmt.Println("X Right Shift 1  = ", x>>1)
	fmt.Println("X Right Shift 2  = ", x>>2)
	fmt.Println("Y Left Shift 1 = ", y<<1)
}

func comparison() {
	a := 45
	b := 12
	fmt.Println("Is A equal to B ? ", a == b)
	fmt.Println("Is A not equal to B ? ", a != b)
	fmt.Println("Is A greater than B ? ", a > b)
	fmt.Println("Is A less than B ? ", a < b)
	fmt.Println("Is A greater than or equal to B ? ", a >= b)
	fmt.Println("Is A less than or equal to B ? ", a <= b)
}

func logical() {
	a := 45
	b := "Something"
	fmt.Println(a > 40 && b == "Something")
	fmt.Println(a < 40 && b == "Something")
	fmt.Println(a < 40 || b == "Something")
	fmt.Println(a < 40 || b != "Something")
	fmt.Println(!(a < 40 || b != "Something"))
}

func assignment() {
	var a int = 100
	b := 20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a += 30
	fmt.Println("a = ", a)
	b -= 5
	fmt.Println("b = ", b)
	a *= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a /= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a %= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
