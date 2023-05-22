package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringIsImmutable1(t *testing.T) {
	str1 := "orange"
	fmt.Printf("str -> %p\n", &str1)
	addr1 := fmt.Sprintf("%p", &str1)

	str2 := str1
	fmt.Printf("str2 -> %p\n", &str2)
	addr2 := fmt.Sprintf("%p", &str2)

	assert.True(t, addr1 != addr2, "String is immutable")
}

func TestStringIsImmutable2(t *testing.T) {
	str := "gopher"
	str_copy := str
	str_copy = "cooper"
	fmt.Println("str = ", str)
	fmt.Println("str_copy = ", str_copy)
}

func TestStringIsImmutable3(t *testing.T) {
	// Character at index cannot be changed in string
	s := "StarWars"
	// s[4] = 'C'
	// s[4] = "C"
	// also won't work
	fmt.Println(s)
}

func TestBoolIsImmutable(t *testing.T) {
	// bool
	boolean := true
	b := boolean
	b = false
	fmt.Println("boolean = ", boolean)
	fmt.Println("b = ", b)
}

func TestPointerIsImmutable(t *testing.T) {
	n := 567
	m := 123

	ptr := &n
	ptr_new := ptr
	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr_new = ", ptr_new)

	ptr_new = &m

	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr_new = ", ptr_new)
}

func TestSliceItemIsMutable(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Printf("S[1] -> %p\n", &s[1])

	p := s
	p[1] = 4
	fmt.Printf("S[1] -> %p\n", &s[1])

	fmt.Println("s =", s)
	fmt.Println("p =", p)

	assert.True(t, &s[1] == &p[1], "Slice item is mutable")
}

func TestArrayItemIsMutable(t *testing.T) {
	a := [3]int{10, 20, 30}
	fmt.Printf("A[1] -> %p\n", &a[1])
	b := a
	b[1] = 40
	fmt.Printf("A[1] -> %p\n", &a[1])

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func TestMapItemIsMutable(t *testing.T) {
	m := map[string]int{"level": 5, "health": 9}
	fmt.Println("m =", m)
	n := m
	n["food"] = 12

	fmt.Println("m =", m)
	fmt.Println("n =", n)
}
