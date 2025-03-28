package main

import (
	"encoding/json"
	"fmt"
)

func main2() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}
