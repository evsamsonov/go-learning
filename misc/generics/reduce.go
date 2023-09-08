package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}

	sum := Reduce(list, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(sum)

	mult := Reduce([]float64{1, 2, 3, 4}, func(a, b float64) float64 {
		return a * b
	}, 1)
	fmt.Println(mult)
}

func Reduce[T any](list []T, accumulator func(a, b T) T, init T) T {
	for _, item := range list {
		init = accumulator(init, item)
	}
	return init
}
