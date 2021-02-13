package __lock_free

import (
	"sync"
	"sync/atomic"
	"testing"
)

func increaseWithLock() {
	var val int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			val++
			mu.Unlock()
		}()
	}
	wg.Wait()
}

func increaseLockFree() {
	var val int64
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for !atomic.CompareAndSwapInt64(&val, val, val+1) {
			}
		}()
	}
	wg.Wait()
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		increaseWithLock()
	}
}

func BenchmarkLockFree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		increaseLockFree()
	}
}
