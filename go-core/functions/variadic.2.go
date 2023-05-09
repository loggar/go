package main

import "fmt"

func main_2() {
	var line string

	line = join_2(",", []string{"Sammy", "Jessica", "Drew", "Jamie"})
	fmt.Println(line)

	line = join_2(",", []string{"Sammy", "Jessica"})
	fmt.Println(line)

	line = join_2(",", []string{"Sammy"})
	fmt.Println(line)
}

func join_2(del string, values []string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}
