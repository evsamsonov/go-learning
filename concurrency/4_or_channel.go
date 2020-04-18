package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}

			// Это выглядит проще, но будет больше рекурсивных вызовов
			//select {
			//case <-channels[0]:
			//case <-channels[1]:
			//case <-or(append(channels[2:], orDone)...):
			//}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	orDone := or(
		sig(time.Minute),
		sig(time.Second*2),
		sig(time.Second*10),
		sig(time.Second),
		sig(time.Hour),
	)
	<-orDone
	fmt.Printf("after %v", time.Since(start))
}
