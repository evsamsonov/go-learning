package main

import (
	"runtime"
	"sync"
	"testing"
)

// see https://habr.com/ru/company/mailru/blog/510200/
//
// $ sysctl -a | grep cacheline
// hw.cachelinesize: 64
//
// $ sysctl hw.l1dcachesize
// hw.l1dcachesize: 32768

const loopLen = 1000

type SimpleStruct struct {
	n int
}

type PaddedStruct struct {
	n int
	_ CacheLinePad
}

const CacheLinePadSize = 64

type CacheLinePad struct {
	_ [CacheLinePadSize]byte
}

func BenchmarkStructureFalseSharing(b *testing.B) {
	structA := SimpleStruct{}
	structB := SimpleStruct{}
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(2)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for j := 0; j < loopLen; j++ {
				structA.n += j
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < loopLen; j++ {
				structB.n += j
			}
		}()
		wg.Wait()
	}
}

func BenchmarkStructureFalseSharingWithCacheLinePad(b *testing.B) {
	structA := SimpleStruct{}
	structB := PaddedStruct{}
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for j := 0; j < loopLen; j++ {
				structA.n += j
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < loopLen; j++ {
				structB.n += j
			}
		}()
		wg.Wait()
	}
}
