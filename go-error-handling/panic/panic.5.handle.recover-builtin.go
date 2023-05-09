package main

import (
	"fmt"
	"log"
)

func main_5() {
	divideByZero_5()
	fmt.Println("we survived dividing by zero!")

}

func divideByZero_5() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide_5(1, 0))
}

func divide_5(a, b int) int {
	return a / b
}
