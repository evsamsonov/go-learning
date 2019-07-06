package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}	// ... - говорит, что размер массива определяется по литералу
	reverse(a[:]) // Передаем как срез
	fmt.Println(a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
