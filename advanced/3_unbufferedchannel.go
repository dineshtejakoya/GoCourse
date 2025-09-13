package advanced

import (
	"fmt"
	"time"
)

//by default a channel in goroutine is an unbuffered channel
//greeting :<- "World"
//Buffered channel means a channel with storage
//Channel Buffering allows channels to hold a limited number of values before blocking the sender
//Buffered channels are useful for managing data flow and controlling concurrency
//The most difference between buffered channel and unbuffered channel is buffered channel allows asynchronous communication,which means buffered channels allows sender to continue working without blocking until the buffer is full
//Buffer channels also help us in load balancing , handling bursts of data without immediate synchronization
//Buffered channels also allow for a better flow control i.e., we can manage the rate of data control between producers and receivers

func main() {
	//unbuffered channels need an immediate receiver, that is why we cannot use a send data into an unbuffered channel inside the main function

	//main thread works faster than the goroutine
	ch := make(chan int)
	//unbuffered channels cannot hold values
	//As soon as the channel receives a value it need to pass the value to a receiving end, we cannot hold those values
	go func() {
		time.Sleep(15 * time.Second)
		fmt.Println("15 seconds finished")
		//ch <- 1
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("5 seconds finished")
		ch <- 1
	}()
	//Receiver will wait for all the goroutines to finish and then it'll throw the error, until then it'll keep the execution flow at halt
	receiver := <-ch
	fmt.Println(receiver)

}

// for unbuffered channels we need to use goroutines
