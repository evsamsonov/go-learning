package main

import "fmt"

func main() {
	queue := make(chan int)
	for i := 0; i < 20; i++ {
		go func() {
			for task := range queue {
				fmt.Printf("Task %d\n", task)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		queue <- i
	}
}
