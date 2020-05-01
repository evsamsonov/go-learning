package main

import (
	"fmt"
	"time"
)

func main() {
	orDone := func(done <-chan struct{}, c <-chan interface{}) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer close(result)
			for {
				select {
				case <-done:
					return
				case val, ok := <-c:
					if !ok {
						return
					}
					select {
					case <-done:
					case result <- val:
					}
				}
			}
		}()
		return result
	}

	done := make(chan struct{})
	defer close(done)

	stream := make(chan interface{})

	go func() {
		for val := range orDone(done, stream) {
			fmt.Println(val)
		}
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			stream <- i
		}
	}()

	time.Sleep(300 * time.Millisecond)
}
