package main

import (
	"testing"
	"time"
)

func doWorkWithPulse(done <-chan interface{}, pulseInterval time.Duration, nums ...int) (chan interface{}, chan int) {
	heartbeat := make(chan interface{}, 1)
	results := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(results)

		time.Sleep(2 * time.Second) // some work
		pulse := time.Tick(pulseInterval)

	numsLoop:
		for _, n := range nums {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case results <- n:
					continue numsLoop
				}
			}
		}
	}()

	return heartbeat, results
}

func Test_doWorkWithPulse(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{1, 3, 7, 8, 2}
	timeout := 2 * time.Second
	heartbeat, results := doWorkWithPulse(done, timeout/2, intSlice...)

	<-heartbeat

	i := 0
	for {
		select {
		case r, ok := <-results:
			if !ok {
				return
			} else if r != intSlice[i] {
				t.Errorf("index %d: expected %d, got %d", i, intSlice[i], r)
			}
			i++
		case <-heartbeat:
		case <-time.After(timeout):
			t.Fatalf("test time out")
		}
	}
}
