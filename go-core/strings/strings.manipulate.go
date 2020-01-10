package main

import (
	"fmt"
	"strings"
)

func manipulate() {
	fmt.Println(strings.Join([]string{"sharks", "crustaceans", "plankton"}, ","))

	balloon := "Sammy has a balloon."
	s := strings.Split(balloon, " ")
	fmt.Println(s)

	data := "  username password     email  date"
	fields := strings.Fields(data)
	fmt.Printf("%q", fields)

}
