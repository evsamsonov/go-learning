package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3)) // При вызове создается массив и срез на весь массив, который передается в функцию

	numbers := []int{4, 5, 6}
	fmt.Println(sum(numbers...)) // Если параметры уже в срезе
}

func sum(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}

	return sum
}
