package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type value struct {
		sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(a *value, b *value) {
		defer wg.Done()
		a.Lock()
		defer a.Unlock()

		time.Sleep(time.Second * 2)
		b.Lock()
		defer b.Unlock()

		fmt.Printf("sum=%d\n", a.value+b.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
