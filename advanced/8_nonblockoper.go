package advanced

import (
	"fmt"
	"time"
)

// Non Blocking operations on channels allow a goroutine to perform a channel operation like send or receive without getting stuck if the channel is not ready
// They help maintain responsiveness and prevent goroutines from getting blocked indefinitely
// it is to prevent deadlock, improve efficiency
func main() {
	//===Non Blocking Receive Operation
	//default value would get printed because we are not sending any value to the channel
	// ch := make(chan int)
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received: ", msg)
	// default:
	// 	fmt.Println("No messages available")
	// }

	//===NON Blocking Send Operation
	//here it will also print default because we don't have a receiver
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent message..")
	// default:
	// 	fmt.Println("Channel is not ready to receive..")
	// }

	//=========Non Blocking Operation in Real Time System
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			//receive value from data channel
			case d := <-data:
				fmt.Println("Data received:", d)
			//if we receive a value on the quit channel, because here it is bool
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)

			}
		}
	}()

	for i := range 5 {
		//send values into data
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}
