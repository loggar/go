package main

import "fmt"

func main_3() {
	var line string

	names := []string{"Sammy", "Jessica", "Drew", "Jamie"}

	line = join_3(",", names...)
	fmt.Println(line)
}

func join_3(del string, values ...string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}
