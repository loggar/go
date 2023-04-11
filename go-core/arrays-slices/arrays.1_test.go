package array1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sum ...
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func TestSum(t *testing.T) {
	array := [...]float64{7.0, 8.5, 9.1}
	s := Sum(&array) // Note the explicit address-of operator

	expected := 24.6
	assert.Equal(t, expected, s, "Sum(%v) should equal %f", array, expected)
}
