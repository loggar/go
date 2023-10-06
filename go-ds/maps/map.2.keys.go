package main

import (
	"fmt"
	"sort"
)

func mapKeys() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)

	fmt.Println(sammy["animal"])
	// returns shark

	fmt.Println(sammy["color"])
	// returns blue

	fmt.Println(sammy["location"])
	// returns ocean

	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}

	keys := []string{}

	for key := range sammy {
		keys = append(keys, key)
	}
	fmt.Printf("%q", keys)

	sort.Strings(keys)

}

func mapValues() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	items := make([]string, len(sammy))

	var i int

	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q", items)
}

func mapLen() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(len(sammy))
}
