package main

import "fmt"

const (
	c1 = 10     // first one must be assigned
	c2          // set to 10
	c3          // set to 10
	c4 = "here" // different type
	c5          // set to "here"
	c6 = iota   // index will be assigned as value (until c8, until assignment is skipped)
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main_2() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}

//  10 10 10 here here 5 6 7 earth earth End
