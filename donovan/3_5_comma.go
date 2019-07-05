package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1212312312345.12"))
	fmt.Println(comma("123412345"))
}

func comma(s string) string {
	if i := strings.LastIndex(s, "."); i >= 0 {
		return comma(s[:i-1]) + s[i:]
	}

	l := len(s)
	if l <= 3 {
		return s
	}

	return comma(s[:l-3]) + "," + s[l-3:]
}
