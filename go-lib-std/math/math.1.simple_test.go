package math_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMathConstants(t *testing.T) {
	fmt.Println("Pi = ", math.Pi)
	fmt.Println("E = ", math.E)
	fmt.Println("Phi = ", math.Phi)
	fmt.Println("Sqrt of 2 = ", math.Sqrt2)
	fmt.Println("Naturla Log 2 = ", math.Ln2)
	fmt.Println("Naturla Log 10 = ", math.Ln10)
}

func TestMathAbs(t *testing.T) {
	a := 45
	b := 100
	diff := a - b
	assert.Equal(t, -55, diff, "diff should equal to %d", -55)

	absolute_diff := math.Abs(float64(a) - float64(b))
	assert.Equal(t, float64(55), absolute_diff, "absolute_diff should equal to %f", 55)
}

func TestMathCasting(t *testing.T) {
	var foo int64 = 77
	fmt.Printf("Type of foo = %T \n", foo)
	fmt.Println("foo = ", int(foo))

	str_foo := strconv.FormatInt(foo, 10)

	fmt.Printf("String Cast: %s\n", str_foo) // Convert to string using base 10
	fmt.Println("Float Cast: ", float64(foo))

	assert.Equal(t, "77", str_foo, "str_foo should equal to %s", "77")
}

func TestMathMinMax(t *testing.T) {
	var a float64 = 120
	var b float64 = 54

	minimum := math.Min(float64(a), float64(b))
	maximum := math.Max(float64(a), float64(b))
	fmt.Printf("Min of %v and %v is %v \n", a, b, minimum)
	fmt.Printf("Max of %v and %v is %v \n", a, b, maximum)

	assert.Equal(t, b, minimum, "minimum should equal to %f", b)
}

func TestMathPowPow10(t *testing.T) {
	var x float64 = 3
	var y float64 = 4
	z := math.Pow(x, y)
	z10 := math.Pow10(int(x))
	fmt.Println("X ^ Y = ", z)
	fmt.Println("10 ^ X = ", z10)

	assert.Equal(t, float64(1000), z10, "10 ^ X should equal to %f", float64(1000))
}

func TestMathSqrt(t *testing.T) {
	var k float64 = 125
	sqrt_of_k := math.Sqrt(k)
	cbrt_of_k := math.Cbrt(k)

	fmt.Printf("Square root of %v = %v \n", k, sqrt_of_k)
	fmt.Printf("Cube root of %v = %v \n", k, cbrt_of_k)

	assert.Equal(t, float64(125), k, "Cube root should equal to %f", float64(125))
}

func TestMathTrunc(t *testing.T) {
	var p float64 = 445.235
	trunc_p := math.Trunc(p)
	fmt.Printf("Truncated value of %v = %v \n", p, trunc_p)
	p = 123.678
	trunc_p = math.Trunc(p)
	fmt.Printf("Truncated value of %v = %v \n", p, trunc_p)

	assert.Equal(t, float64(123), trunc_p, "Truncated value should equal to %f", float64(123))
}

func TestMathCeil(t *testing.T) {
	var c float64 = 33.25
	ceil_c := math.Ceil(c)
	fmt.Printf("Ceiled value of %v = %v \n", c, ceil_c)
	c = 134.78
	ceil_c = math.Ceil(c)
	fmt.Printf("Ceiled value of %v = %v \n", c, ceil_c)

	assert.Equal(t, float64(134.78), c, "Ceiled value should equal to %f", float64(134.78))
}

func TestMathSinCosTan(t *testing.T) {
	// basic trigonometric functions
	var x float64 = math.Pi / 2
	sinx := math.Sin(x)
	cosx := math.Cos(x)
	tanx := math.Tan(x)
	fmt.Printf("Sin(%v) = %v \n", x, sinx)
	fmt.Printf("Cos(%v) = %v \n", x, cosx)
	fmt.Printf("Tan(%v) = %v \n", x, tanx)

	// hyperbolic trigonometric functions
	var h float64 = math.Pi / 2
	sinh := math.Sinh(h)
	cosh := math.Cosh(h)
	tanh := math.Tanh(h)
	fmt.Printf("Sinh(%v) = %v \n", h, sinh)
	fmt.Printf("Cosh(%v) = %v \n", h, cosh)
	fmt.Printf("Tanh(%v) = %v \n", h, tanh)

	// Inverse Trigonometric functions
	var y float64 = -1
	arc_sin := math.Asin(y) // -pi/2 radians or 90 degrees
	arc_cos := math.Acos(y) // pi randians or 180 degrees
	arc_tan := math.Atan(y)
	fmt.Printf("Sin^-1(%v) = %v \n", y, arc_sin)
	fmt.Printf("Cos^-1(%v) = %v \n", y, arc_cos)
	fmt.Printf("Tan^-1(%v) = %v \n", y, arc_tan)

	assert.Equal(t, float64(-0.7853981633974483), arc_tan, "Tan^-1(%v) should equal to %f", y, float64(-0.7853981633974483))
}

func TestMathExp(t *testing.T) {
	// exponential function
	var x float64 = 2
	y := math.Exp(x)
	fmt.Println("e^x = ", y)
	var n float64 = 3.5
	y = math.Exp2(n)
	fmt.Println("2^n = ", y)

	// Logarithmic function
	y = math.Log(x)
	fmt.Println("natural log x = ", y)

	n = 128
	y = math.Log2(n)
	fmt.Println("Log2 of 128 = ", y)

	assert.Equal(t, float64(7), y, "Log2 of 128 should equal to %f", float64(7))
}
