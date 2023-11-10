package constraint

type Number interface {
	int | float32 | float64
	IsEven() bool
}
