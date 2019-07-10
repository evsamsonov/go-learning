package main

import "fmt"

func main() {
	var slice []int

	for i := 0; i < 10; i++ {
		slice = appendInt(slice, i)
		fmt.Println(slice, cap(slice))
	}

}

func appendInt(slice []int, elem int) []int {
	var newSlice []int

	newLen := len(slice) + 1
	if newLen <= cap(slice) {
		// Есть место, расширяем срез
		newSlice = slice[:newLen]
	} else {
		// Места нет, нужен массив больше
		newCap := newLen
		if newCap < 2*len(slice) {
			newCap = 2 * len(slice)
		}

		newSlice = make([]int, newLen, newCap)
		copy(newSlice, slice) // Встроенная функция для копирования
	}

	newSlice[len(slice)] = elem
	return newSlice
}
