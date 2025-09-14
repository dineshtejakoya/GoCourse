package concurrency

import (
	"fmt"
	"sync"
)

//=======here in increment() function uncomment c.mu.Lock() and defer c.mu.Unlock() and run  >>go run -race 2_rc.go in linux

type counter struct {
	//sync.Mutex is used to ensure that only one goroutine can access the value at a time,where multiple goroutines might try to modify value simultaneously
	mu sync.Mutex
	//count that multiple goroutines will try to modify
	count int
}

// adding few methods of this type
func (c *counter) increment() {
	//c.mu.Lock()
	//we need to unlock at the end of the function
	//defer c.mu.Unlock()
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
