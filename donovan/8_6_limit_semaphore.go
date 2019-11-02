package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		doSomething()
	}
}

// Ограничивающий семафор на 20 горутин
var limitSemaphore = make(chan struct{}, 20)
func doSomething() {
	fmt.Println(cap(limitSemaphore))

	limitSemaphore <- struct{}{}

	// Что-то делаем

	<-limitSemaphore
}
