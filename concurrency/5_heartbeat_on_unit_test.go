package main

import (
	"testing"
	"time"
)

func doWork(done <-chan interface{}, nums ...int) (chan interface{}, chan int) {
	heartbeat := make(chan interface{}, 1)
	results := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(results)

		time.Sleep(2 * time.Second) // some work

		for _, num := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case results <- num:
			}
		}

	}()

	return heartbeat, results
}

func Test_doWork(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{1, 3, 7, 8, 2}
	heartbeat, results := doWork(done, intSlice...)

	<-heartbeat

	for i, expected := range intSlice {
		select {
		case r := <-results:
			if expected != r {
				t.Errorf("index %d: expected %d, got %d", i, expected, r)
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("test time out")
		}
	}
}
