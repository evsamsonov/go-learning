package main

import "fmt"

func main() {
	generator := func(done <-chan struct{}) <-chan int {
		result := make(chan int)
		go func() {
			for i := 0; i < 1000; i++ {
				select {
				case <-done:
					return
				case result <- i:
				}
			}
			close(result)
		}()
		return result
	}

	multiply := func(done <-chan struct{}, values <-chan int) <-chan int {
		result := make(chan int)
		go func() {
			for {
				select {
				case <-done:
					return
				case val, ok := <-values:
					if !ok {
						close(result)
						return
					}
					result <- val * val
				}
			}
		}()
		return result
	}

	done := make(chan struct{})
	defer close(done)

	values := multiply(done, generator(done))
	for val := range values {
		fmt.Println(val)
	}
}
