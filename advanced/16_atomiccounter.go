package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Atomic Counter is a type of counter used in concurrent programming to manage or count values in a thread safe manner without the need for explicit locking
Performance - Atomic Operations are often faster than mutexes due to lower overhead
Simplicity - it provides a straight forward way to manage counters in concurrent environments, easier to implement in applications
concurrency - ensures that counter updates are safe and consistent across multiple goroutines

Atomic Operations in Go:
Window Duration
Request Counting
Reset Mechanism

sync/atomic package:
atomic.AddInt32/atomic.AddInt64
atomic.LoadInt32/atomic.LoadInt64
atomic.StoreInt32/atomic.StoreInt64
atomic.CompareAndSwapInt32/atomic.CompareAndSwapInt64

Atomic means Indivisible,Uninterruptible(i.e., it runs in a single step without interruption)
*/

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	//1 is delta which will be added to initial value
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{}
	//value := 0

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				//value++
			}
		}()
	}

	wg.Wait()
	//fmt.Printf("Final Counter Value: %d\n", value)
	fmt.Printf("Final Counter Value: %d\n", counter.getValue())
}

/*
If you don't use an atomic counter and instead access a simple variable from multiple goroutines to perform increments
encounter several issues related to data races and incorrect results
Data Races: A data race occurs when two or more goroutines access the same variable concurrently and atlease one of those accesses is a write, since there is no
synchronization operation in place, the result of these operations is unpredictable
*/
