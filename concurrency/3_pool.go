package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Created new instances")
			return struct{}{}
		},
	}

	instance := myPool.Get() // call New
	myPool.Put(instance)     // return to pull
	myPool.Get()             // dont call New
	myPool.Get()             // call New
}
