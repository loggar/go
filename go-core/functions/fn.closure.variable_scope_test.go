package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func inrementer() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func TestClosureScope(t *testing.T) {
	c := inrementer()
	v := c()
	assert.Equal(t, 1, v, "closure, variable scope")

	v = c()
	assert.Equal(t, 2, v, "closure, variable scope")

	v = c()
	assert.Equal(t, 3, v, "closure, variable scope")

	d := inrementer()
	v = d()
	assert.Equal(t, 1, v, "closure, variable scope")
}

func factorial() func() int {
	fact, n := 1, 1
	return func() int {
		fact = fact * n
		n += 1
		return fact
	}
}

func TestFactorial(t *testing.T) {
	f := factorial()
	v := f()
	assert.Equal(t, 1, v, "closure, factorial")

	v = f()
	assert.Equal(t, 2, v, "closure, factorial")

	v = f()
	assert.Equal(t, 6, v, "closure, factorial")
}
