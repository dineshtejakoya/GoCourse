package advanced

import (
	"fmt"
	"time"
)

// Channels are a way for goroutines to communicate with each other and synchronize their execution
// they provide a means to send and receive values between goroutines facilitating data exchange and coordination
// we use make function to create a channel
// variable := make(chan type)
func main() {
	//channel can be of any type(a list of structs as well)
	//we cannot make our channel work inside our function, we have to create goroutine inorder to receive the channel
	greeting := make(chan string)
	greetString := "Hello"

	//here it tries to send value to a channel without having a goroutine ready
	//without it will causes a deadlock,because channels in go are blocking
	//goroutines are non blocking, they are extracted away from the main thread, channels in go are blocking
	//greeting <- greetString

	//here greeting channel is an open channel receiving continuous values
	//suppose a weather forecast receives data every 5 minutes we use in this case
	go func() {
		greeting <- greetString
		//now if we pass one more channel, it won't be received with one receiver
		//we need to initialize the receiver again
		//greeting <- "World"
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// }()

	//receiver is receiving value from the greeting channel
	// go func() {
	// 	receiver := <-greeting
	// 	//2.
	// 	fmt.Println(receiver)
	// }()
	receiver := <-greeting
	fmt.Println(receiver)

	receiver = <-greeting
	fmt.Println(receiver)

	//for _ = range 5 {
	for range 5 {
		revr := <-greeting
		fmt.Println(revr)
	}

	fmt.Println("End of program")

	time.Sleep(1 * time.Second)
}

//main function is a goroutine,
//Goroutine = Thread simple..,
//similar to goroutine channesl are also managed by goroutine

//receiving from a channel is not blocking always,
//if there is no value to receive then it will wait for the value to received and next line will not be executed until it receives the value
//Receiver is only blocking until the time it hasn't received a value, as soon as it receives a value it won't block
//to prevent that place the steps in another goroutine and wait for it to finish
