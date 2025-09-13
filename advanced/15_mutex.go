package advanced

import (
	"fmt"
	"sync"
)

/*
mutex in short for mutual exclusion
it is a synchronization primitive, used to prevent multiple goroutines from simultaneously accessing shared resources or executing critical sections of code
it ensures that only one goroutine can hold a mutex at a time, thus avoiding race conditions and data corruptions

Why we need Mutexes?
Data Integrity -- we need to protect our resources from concurrent modifications
Synchronization -- Mutexes coordinate access to critical section of code

Basic Operations:
Lock() - acquires mutex and protects it from getting it accessed by other goroutines,
	If it is not able to access the goroutine it is blocking in nature if another goroutine holds the target
Unlock() - releases the mutex allowing other mutexes to acquire it
TryLock() - returns a boolean value, it attempts to aquire the boolean value without blocking it and returns a boolean value if it is able acquire the mutex or not
*/

type counter struct {
	//sync.Mutex is used to ensure that only one goroutine can access the value at a time,where multiple goroutines might try to modify value simultaneously
	mu sync.Mutex
	//count that multiple goroutines will try to modify
	count int
}

// adding few methods of this type
func (c *counter) increment() {
	c.mu.Lock()
	//we need to unlock at the end of the function
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	//it locks the counter before accessing value and unlocks it afterward ensuring thread safe access
	//simply , lock it access the value and unlock it
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	//it creates a new instance of the counter structs and returns the pointer to it
	//shorthand for new counter,
	counter := &counter{}

	numGoroutines := 10

	//wg.Add(numGoroutines)

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				//if we access direct counter value with out mutex like below,it would be unreliable
				//counter.count++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())
}

/*
When we are not using mutex our count value is initialized to zero
4 goroutines are receiving value zero,queued up all gave value back as 1 instead of 1
goroutines are getting wasted
*/
/*
when we use mutexes ,
4 goroutines in place
*/
