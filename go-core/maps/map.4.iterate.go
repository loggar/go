package maps

import (
	"fmt"
)

func IterateMapForLoop() {
	is_prime := map[int]bool{
		7:  true,
		9:  false,
		13: true,
		15: false,
		16: false,
	}

	for key, value := range is_prime {
		fmt.Printf("%d -> %t\n", key, value)
	}
}

func IterateMapRange() {
	is_prime := map[int]bool{
		7:  true,
		9:  false,
		13: true,
		15: false,
		16: false,
	}

	for key, _ := range is_prime {
		fmt.Printf("Key : %d\n", key)
	}

	for _, value := range is_prime {
		fmt.Printf("Value: %t\n", value)
	}
}
