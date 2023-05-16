package math_test

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathBits(t *testing.T) {
	s, c := bits.Add(0, 9, 1)
	fmt.Printf("Sum = %d \nCarry = %d \n", s, c)

	// (45) in decimal = (1 0 1 1 0 1) in binary
	var n uint = 45
	length := bits.Len(n)
	ones_in_45 := bits.OnesCount(n)
	fmt.Printf("Minimum bits required to represent 45 = %d \n", length)
	fmt.Printf("Set Bits in 45 = %d \n", ones_in_45)

	assert.Equal(t, 4, ones_in_45, "Bits in 45 is unexpected value")
}
