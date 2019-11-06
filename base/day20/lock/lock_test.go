package lock

import (
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {

	count := 0
	for i := 0; i < 10000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(5 * time.Second)
	t.Logf("count = %v", count)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	count := 0
	for i := 0; i < 10000; i++ {
		go func() {
			defer func() {
				mut.Unlock() // 使用 defer 释放锁
			}()
			mut.Lock() // 加锁
			count++
		}()
	}
	time.Sleep(5 * time.Second)
	t.Logf("count = %v", count)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock() // 使用 defer 释放锁
			}()
			mut.Lock() // 加锁
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("count = %v", count)
}

func TestCounterRWLock(t *testing.T) {
	var rwmut sync.RWMutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				rwmut.Unlock() // 使用 defer 释放锁
			}()
			rwmut.Lock() // 加锁
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("count = %v", count)
}
