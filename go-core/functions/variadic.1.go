package main

import "fmt"

func main() {
	sayHello()
	sayHello("Sammy")
	sayHello("Sammy", "Jessica", "Drew", "Jamie")
}

func sayHello(names ...string) {
	if len(names) == 0 {
		fmt.Println("nobody to greet")
		return
	}
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
