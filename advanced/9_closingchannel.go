package advanced

import "fmt"

/*
Why Close Channels?
We close channels to signal completion, it indicates that no more data will be sent on the channel which helps goroutines that receiving data know when to stop waiting
It also prevents resource leaks, closing channels ensures that resources associated with the channels are properly cleaned up
*/
//After a channel is closed no more values can be sent to it, however we can receive from a closed channel if it is a buffered channel
//A buffered channel may have some values stored in the buffer and those values can be received

//Simple Closing channel example
// func main() {
// 	ch := make(chan int)

// 	//sending values to channel
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	//Receiving values from channel
// 	for val := range ch {
// 		fmt.Println(val)
// 	}

// }

//Receiving from a closed channel
// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel Closed")
// 		return
// 	}
// 	fmt.Println(val)
// }

//Range Over Closed Channel
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 	}
// }

/*============BEST Practices==============
Close channels only from the sender
Avoid closing channel more than once, if does it will create runtime panic in action
*/

//Closing channel more than once
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		close(ch)
// 		close(ch)
// 	}()
// 	time.Sleep(1 * time.Second)
// }

/*=========COMMON Patterns for closing channels==========
Pipeline Pattern
Worker Pool Pattern
*/

//send only channel
func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

//in ->Receive only channel
//out -> send only channel
func filter(in <-chan int, out chan<- int) {
	//looping over in channel
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	//as we are using out as send only channel, it is responsibility of sender to close it
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}

//Another pattern to close channel is worker pool
