package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(s, 2))
}

// Удаление элемента с сохранением порядка
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	// slice[i] = slice[len(slice) - 1] 	// без сохранение порядка
	return slice[:len(slice)-1]
}
