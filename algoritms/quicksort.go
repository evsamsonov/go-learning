package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	array := []int{2323, 432, 12, 534, 76, 0, 3, 500, 1}
	fmt.Println(quicksort(array))
}

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	pivotIndex := rand.Intn(len(array))
	pivot := array[pivotIndex]

	var less []int
	var greater []int
	for index, num := range array {
		if pivotIndex == index {
			continue;
		}

		if num < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = append(quicksort(less), pivot)
	greater = quicksort(greater)
	return append(less, greater...)
}
