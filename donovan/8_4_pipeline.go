package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}

		close(naturals)
	}()

	go func() {
		for i := range naturals {
			squares <- i * i
		}

		close(squares)
	}()

	for i := range squares {
		fmt.Println(i)
	}
}
