package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	PageNumber int      `json:"page"`
	Fruits     []string `json:"fruits"`
}

func testMain() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.PageNumber)
}

//=> 1
