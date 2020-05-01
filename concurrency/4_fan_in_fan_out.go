package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	primeFinder := func(done <-chan struct{}, intStream <-chan int) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for integer := range intStream {
				integer -= 1
				prime := true
				for div := integer - 1; div > 1; div-- {
					if integer%div == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}

	done := make(chan struct{})
	defer close(done)

	randInt := func(done <-chan struct{}) <-chan int {
		result := make(chan int)
		go func() {
			for {
				select {
				case <-done:
				case result <- rand.Intn(50000000):
				}
			}
		}()
		return result
	}

	fanIn := func(done <-chan struct{}, channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
				case multiplexStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexStream)
		}()

		return multiplexStream
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

	randIntStream := randInt(done)

	// Fan Out?
	numFinders := runtime.NumCPU()
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Println(prime)
	}
}
