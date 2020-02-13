package main

import "fmt"

func main() {
	sammyShark := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	for key, value := range sammyShark {
		fmt.Println(key + ": " + value)
	}
}
