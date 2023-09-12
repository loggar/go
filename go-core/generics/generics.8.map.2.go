package main

import (
	"fmt"
)

func GetOrSetValue[K comparable, V any](m *map[K]V, key K, defaultVal V) V {
	// reference the original map
	ref := *m
	if v, ok := ref[key]; ok {
		return v
	} else {
		//mutate the original map
		ref[key] = defaultVal

		return defaultVal
	}
}

func usage_get_value_pass_by_ref() {
	serverStats := map[string]int{
		"port":      8000,
		"pings":     47,
		"status":    1,
		"endpoints": 13,
	}
	fmt.Println(serverStats)
	v := GetOrSetValue(&serverStats, "cpu", 4)
	fmt.Println(v)
	fmt.Println(serverStats)
}
