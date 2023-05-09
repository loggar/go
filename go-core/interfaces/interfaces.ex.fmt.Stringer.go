package main

import (
	"fmt"
	"strings"
)

type Ocean struct {
	Creatures []string
}

func (o Ocean) String() string {
	return strings.Join(o.Creatures, ", ")
}

func log(header string, s fmt.Stringer) {
	fmt.Println(header, ":", s)
}

func test_stringer() {
	o := Ocean{
		Creatures: []string{
			"sea urchin",
			"lobster",
			"shark",
		},
	}
	log("ocean contains", o)
}
