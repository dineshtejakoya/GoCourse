package advanced

import "fmt"

//Buffered channels allow channels to hold a limited number of values before blocking the sender
//useful fo rmanaging data flow and controlling concurrency
//we are allowing channesl to store values, earlier in unbuffered channels we saw they need an immediate outflow of the channel
//buffered channels do not require an immediate receiver
//they allow senders to continue working without blocking until the buffer is full
//they do not require an immediate receiver

//other than that we also use buffered channels when we are handling bursts of data without immediate synchronization, means the channel is receiving values and sending values somewhere we don't need that because we are storing the values of buffer.
//water pipe they will fled out when we open, but pipe has a capacity that can be decided by us during declaration
func main() {
	//make(chan Type,capacity)
	ch := make(chan int, 2)
	//principles on buffered channels
	//buffered channels block on send only if the buffer is full
	//buffered channels block on receive only if the buffer is empty

	ch <- 1
	//it does not need an immediate receiver and our execution moves to next line
	ch <- 2
	//dead lock will occur as capacity is only 2
	//ch <- 3
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	//now if try to insert third value there wouldn't be any error
	ch <- 3
	fmt.Println("Buffered Channels")
}
