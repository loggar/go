package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPerson struct {
	name string
	age  int
}

func TestPointerClone(t *testing.T) {
	// make a struct pointer variable a
	a := &testPerson{name: "Alice", age: 30}

	// make b := a (simple pointer copy)
	b := a

	fmt.Printf("pointer a: %p, value: %+v\n", a, *a)
	fmt.Printf("pointer b: %p, value: %+v\n", b, *b)

	// assert pointer a and b are the same
	assert.Same(t, a, b, "a and b should reference the same variable")

	v1 := *a
	v2 := *b

	// assert a and b are equal (by value)
	assert.Equal(t, v1, v2, "a and b should be equal by value")

	a.age = 31

	assert.Equal(t, a.age, 31, "a.age should be updated to 31")
	assert.Equal(t, v1.age, 30, "v1.age should remain 30")

	// now make c as a clone of a
	c := *a // dereference a to get value, assign to c (new variable)

	fmt.Printf("pointer a: %p, value: %+v\n", a, *a)
	fmt.Printf("pointer c: %p, value: %+v\n", &c, c)

	// assert pointer a and c are not the same
	assert.NotSame(t, a, &c, "a and c should not reference the same variable")

	v3 := c

	// assert a and c are equal (by value)
	assert.Equal(t, *a, v3, "a and c should be equal by value")

	a.age = 32

	assert.Equal(t, a.age, 32, "a.age should be updated to 32")
	assert.Equal(t, v3.age, 31, "v3.age should remain 31")
}
