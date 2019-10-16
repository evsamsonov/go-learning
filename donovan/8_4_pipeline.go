package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counters(naturals)
	go squarer(naturals, squares)

	for i := range squares {
		fmt.Println(i)
	}
}

func counters(naturals chan<- int) {
	for i := 0; i < 100; i++ {
		naturals <- i
	}

	close(naturals)
}

func squarer(naturals <-chan int, squares chan<- int) {
	for i := range naturals {
		squares <- i * i
	}

	close(squares)
}
