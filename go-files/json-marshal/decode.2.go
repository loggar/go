package main

import (
	"encoding/json"
	"fmt"
)

type responseType struct {
	Name *string `json:"name"`
}

func main() {
	str := `{"a": 1, "b": ["apple", "peach"], "name": null}`
	res := responseType{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Name)
}

//=> <nil>
