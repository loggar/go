package main

import (
	"fmt"
)

func GetValue[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultVal
}

func usage_get_value_pass_by_value() {
	serverStats := map[string]int{
		"port":      8000,
		"pings":     47,
		"status":    1,
		"endpoints": 13,
	}
	v := GetValue(serverStats, "status", -1)
	fmt.Println(v)
	v = GetValue(serverStats, "cpu", 4)
	fmt.Println(v)
}
