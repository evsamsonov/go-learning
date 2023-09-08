package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	keys := MapKeys[string, int](m)
	fmt.Println(keys)
	// or
	keys = MapKeys(m)
	fmt.Println(keys)
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	var s []K
	for k := range m {
		s = append(s, k)
	}
	return s
}
