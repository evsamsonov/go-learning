package main

import (
	"context"
	"fmt"
	"sync"
)

type IntSharedVal struct {
	val   int
	write chan int
	read  chan int
}

func (v *IntSharedVal) Init(ctx context.Context) {
	v.write = make(chan int)
	v.read = make(chan int)
	go func() {
		for {
			select {
			case val := <-v.write:
				v.val = val
				fmt.Printf("write %d\n", val)
			case v.read <- v.val:
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (v *IntSharedVal) Set(val int) {
	v.write <- val
}

func (v *IntSharedVal) Get() int {
	return <-v.read
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	intSharedVal := IntSharedVal{}
	intSharedVal.Init(ctx)

	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			intSharedVal.Set(i)
			fmt.Printf("val: %d\n", intSharedVal.Get())
		}(i)
	}

	wg.Wait()
	cancel()
}
