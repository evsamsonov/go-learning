package main

import "fmt"

func main() {
	fmt.Println(equalMap(map[string]int{"a": 1}, map[string]int{"a": 1}))
	fmt.Println(equalMap(map[string]int{"a": 1}, map[string]int{"b": 1}))
}

func equalMap(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for ak, av := range a {
		if bv, ok := b[ak]; !ok || bv != av {
			return false
		}
	}

	return true
}
