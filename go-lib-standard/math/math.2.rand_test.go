package math_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathRand(t *testing.T) {
	// random integer generation
	x := rand.Int()
	fmt.Println(x)

	// random number generation till range
	for i := 0; i < 100; i++ {
		y := rand.Intn(10)
		b := y >= 0 && y < 10
		assert.True(t, b, "y > 0 && y < 10")
	}
}
