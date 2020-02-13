package main

import "fmt"

func main() {
	grade := 92
	if grade >= 65 {
		fmt.Print("Passing grade of: ")

		if grade >= 90 {
			fmt.Println("A")

		} else if grade >= 80 {
			fmt.Println("B")

		} else if grade >= 70 {
			fmt.Println("C")

		} else if grade >= 65 {
			fmt.Println("D")
		}

	} else {
		fmt.Println("Failing grade")
	}
}
