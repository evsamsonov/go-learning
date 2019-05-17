package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
}

func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return factorial(x - 1) * x
	}
}