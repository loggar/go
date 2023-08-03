package main

import (
	"fmt"
	"strings"
)

func get_csv(index int, str string, t func(csv string) []string) [][]string {
	s := t(str)
	var res [][]string
	for i := 1; i < len(s); i += 2 {
		var data []string
		data = append(data, s[i-1], s[i])
		res = append(res, data)
	}
	return res
}

func get_csv_caller() {
	csv_slice := func(csv string) []string {
		return strings.Split(csv, ",")
	}
	csv_data := get_csv(2, "kevin,21,john,33,george,24", csv_slice)
	fmt.Println(csv_data)
	for _, s := range csv_data {
		fmt.Printf("%s - %s\n", s[0], s[1])
	}
}
