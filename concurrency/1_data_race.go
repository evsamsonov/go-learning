package main

import (
	"fmt"
	"sync"
)

func main() {
	var num int
	var mu sync.Mutex
	go func() {
		mu.Lock()
		num++
		mu.Unlock()
	}()

	mu.Lock()
	if num == 0 {
		fmt.Printf("the value is 0\n")
	} else {
		fmt.Printf("the value is %v\n", num)
	}
	mu.Unlock()
}
