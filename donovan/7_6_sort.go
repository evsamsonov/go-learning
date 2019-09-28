package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var stringSlice = []string{
		"a",
		"z",
		"c",
		"g",
		"b",
	}

	sort.Sort(StringSlice(stringSlice))
	fmt.Println(stringSlice)

	sort.Sort(sort.Reverse(StringSlice(stringSlice)))
	fmt.Println(stringSlice)
}
