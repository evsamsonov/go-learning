package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCount := 0
loop:
	for {
		select {
		case <-done:
			break loop // allow to break from for
		default:
		}

		workCount++
		time.Sleep(time.Second)
	}

	fmt.Printf("workCount = %d", workCount)
}
