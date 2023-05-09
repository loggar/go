package main

import "fmt"

func main_3() {
	c := struct {
		Name string
		Type string
	}{
		Name: "Sammy",
		Type: "Shark",
	}
	fmt.Println(c.Name, "the", c.Type)
}
