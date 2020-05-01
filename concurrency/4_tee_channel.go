package main

import "fmt"

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

	tee := func(done <-chan struct{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for val := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					case out1 <- val:
						// Запись в нулевой канал вызывает deadlock
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}

	done := make(chan struct{})
	defer close(done)

	inStream := make(chan interface{})
	out1, out2 := tee(done, inStream)

	go func() {
		for i := 0; i < 10; i++ {
			inStream <- i
		}
		close(inStream)
	}()

	for {
		select {
		case val, ok := <-out1:
			if !ok {
				return
			}
			fmt.Printf("out1: %v\n", val)
		case val, ok := <-out2:
			if !ok {
				return
			}
			fmt.Printf("out1: %v\n", val)
		}
	}
}
