package main

import "fmt"

func main() {
	array := []int{2323, 432, 12, 534, 76, 0, 3, 500}
	fmt.Println(sort(array))
}

func sort(array []int) []int {
	newArray := make([]int, len(array))

	for i := 0; i < len(newArray); i++ {
		minIndex := findMinIndex(array)
		newArray[i] = array[minIndex]
		array = append(array[:minIndex], array[minIndex+1:]...)
	}

	return newArray
}

func findMinIndex(array []int) int {
	var min int
	minIndex := 0

	// или for i, value := range array {
	for i := 0; i < len(array); i++ {
		if i == 0 || array[i] < min {
			min = array[i]
			minIndex = i
		}
	}

	return minIndex
}
