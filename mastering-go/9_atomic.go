package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var val int64

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go testAtomic(&wg)
	}

	wg.Wait()
	fmt.Println(val)
}

func testAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&val, 1)
	}
}
