package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			once.Do(increment)
			wg.Done()
		}
	}()

	wg.Wait()

	fmt.Println(count)
}
