package main

import "fmt"

func main() {
	fmt.Println(Ternary(true, 1, 2))
	fmt.Println(Ternary(false, "abc", "xzy"))
}

func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
