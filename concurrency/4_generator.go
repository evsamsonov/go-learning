package main

import (
	"fmt"
)

func main() {
	repeat := func(done <-chan struct{}, values ...interface{}) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer close(result)
			for {
				for _, val := range values {
					select {
					case <-done:
						return
					case result <- val:
					}
				}
			}
		}()
		return result
	}

	// Позволяет выбрать ограниченное кол-во значений из генератора
	take := func(done <-chan struct{}, valuesChan <-chan interface{}, num int) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer close(result)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case result <- <-valuesChan:
				}
			}
		}()
		return result
	}

	toInt := func(done <-chan struct{}, valuesChan <-chan interface{}) <-chan int {
		result := make(chan int)
		go func() {
			defer close(result)
			for val := range valuesChan {
				select {
				case <-done:
					return
				case result <- val.(int):
				}
			}
		}()
		return result
	}

	done := make(chan struct{})
	defer close(done)

	for val := range toInt(done, take(done, repeat(done, 1), 10)) {
		fmt.Println(val)
	}

}
