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

	bridge := func(done <-chan struct{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case <-done:
					return
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}
					stream = maybeStream
				}
				for val := range orDone(done, stream) {
					select {
					case <-done:
					case valStream <- val:
					}
				}
			}
		}()
		return valStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for val := range bridge(nil, genVals()) {
		fmt.Println(val)
	}

}
