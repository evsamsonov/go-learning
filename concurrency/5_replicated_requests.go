package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, id int, wg *sync.WaitGroup, result chan<- int) {
		started := time.Now()
		defer wg.Done()

		simulateLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
		select {
		case <-done:
		case <-time.After(simulateLoadTime):
		}

		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)
		if took < simulateLoadTime {
			took = simulateLoadTime
		}
		fmt.Printf("%v took %v\n", id, took)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	fmt.Printf("firstReturned=%v\n", firstReturned)
}
