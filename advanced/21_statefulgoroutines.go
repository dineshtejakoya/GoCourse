package advanced

import (
	"fmt"
	"time"
)

//stateful goroutines is a goroutines that maintains and updates its own internal state across multiple invocations or interactions
//it keeps track of state information

type StatefulWorker struct {
	count int
	ch    chan int
}

func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)

			}
		}
	}()
}

// send value to the channel
func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func main() {
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}

	stWorker.Start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}

}
