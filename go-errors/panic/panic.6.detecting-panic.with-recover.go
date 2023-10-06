package main

import (
	"fmt"
	"log"
)

func main_6() {
	divideByZero_6()
	fmt.Println("we survived dividing by zero!")

}

func divideByZero_6() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide_6(1, 0))
}

func divide_6(a, b int) int {
	if b == 0 {
		panic(nil)
	}
	return a / b
}
