package constraint

import "fmt"

type CustomType int16

func printValueSimpleUnion[T int16 | CustomType](value T) {
	fmt.Println(value)
}

func printValueUnderlyingType[T ~int16](value T) {
	fmt.Printf("Value %d", value)
}
