package main

import (
	"fmt"
	"strconv"
)

func main_2() {
	var i uint64 = 45
	var str string
	str = strconv.FormatUint(i, 10)

	fmt.Println(str)
}
