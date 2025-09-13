package advanced

import (
	"fmt"
	"time"
)

/*
Process of handling multiple channel operations simultaneously allowing a goroutine to wait on multiple channel operations and react to whichever operation is ready firt
The select statement in go facilitates multiplexing by allowing a goroutine to wait on multiple channel operations
multiplexing is like a switch
Multiplexing manages multiple concurrent operations within a single goroutines
Multiplexing efficiently handles operations that might block without locking up resources
We get to implement timeout and cancellation mechanisms
*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//As above channels are unbuffered channels hence we cannot send values in the main thread, we have to create a goroutine
	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 1
	}()

	time.Sleep(time.Second)
	//=====SELECT Statement
	// select {
	// //At one point only one case will be executed out of all these cases
	// case msg := <-ch1:
	// 	//first channel that becomes ready will execute its corresponding case
	// 	fmt.Println("Received from ch1:", msg)
	// case msg := <-ch2:
	// 	fmt.Println("Received from ch2:", msg)
	// default:
	// 	fmt.Println("No channels ready..")
	// }

	//we are trying to receive but we don't have any sender
	//without switch we would receive a deadlock, now with switch it has been handled gracefully
	//earlier msg:= <-ch1 is a statement with switch it became a condition

	//If we want to receive two messages place it in a for loop depending on number of channels
	for range 2 {
		select {
		//At one point only one case will be executed out of all these cases
		case msg := <-ch1:
			//first channel that becomes ready will execute its corresponding case
			fmt.Println("Received from ch11:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch22:", msg)
		default:
			fmt.Println("No channels ready..")
		}
	}

	fmt.Println("End of Program...")

}
