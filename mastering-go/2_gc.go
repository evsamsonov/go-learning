package main

import (
	"fmt"
	"runtime"
)

var mem runtime.MemStats

func main() {
	printStats()
	// mem.Alloc:  81496
	// mem.TotalAlloc:  81496
	// mem.HeapAlloc:  81496
	// mem.NumGC:  0
	// -----

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		_ = s
	}
	printStats()
	// mem.Alloc:  72288
	// mem.TotalAlloc:  500125824
	// mem.HeapAlloc:  72288
	// mem.NumGC:  10
	// -----
}

func printStats() {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc: ", mem.Alloc)
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc)
	fmt.Println("mem.NumGC: ", mem.NumGC)
	fmt.Println("-----")
}
