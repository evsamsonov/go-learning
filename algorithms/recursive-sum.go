package main

import "fmt"

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5, 10000}))
}

func sum(array []int) int {
	if len(array) == 0 {
		return 0
	}

	return array[0] + sum(array[1:])
}
