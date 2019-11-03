package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		// Если несколько case срабатывает, то выбирается случайный
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
