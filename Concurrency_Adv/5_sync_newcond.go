package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// sync.NewCond is a function in Go's sync package that create a new condition variable
// A condition variable is a synchronization primitive that allows goroutines to wait for certain conditions to be met while holding a lock,
// It is used to signal one or more goroutines that some condition has changed
const bufferSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	//sync.NewCond creates a new conditionn variable associated with the buffers mutex which it takes as an argument
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		//pauses the goroutine until a signal is received indicating space is available
		b.cond.Wait()
		//releases the mutex temporarily and the for loop keeps on checking is equal to buffer size or not
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item, time.Now())
	//it wakes up the other goroutine
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == 0 {
		b.cond.Wait()
		//This function stops doing anything and waits for other functions to append the slice
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item, time.Now())
	b.cond.Signal()
	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	//we can give common value for range 10 in both producer and consumer
	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
