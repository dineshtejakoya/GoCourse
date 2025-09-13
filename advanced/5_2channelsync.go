package advanced

//string('0'+i) explore
//Are we closing sender or receiver?

import (
	"fmt"
	"time"
)

// Synchronizing Data Exchange
func main() {
	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			//data <- "hello " + string(i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()

	//close(data) //channel closed before goroutine could send a value to the channel

	//Alternate way to create receiver for a channel
	//if we have a channel and it is continuously sending data, we can loop over that cahnnel and it'll create receiver
	//simply here we are ranging over channel, we are not creating receivers with arrow operators this time

	//here order of channels is maintained properly, earlier  it was not predicted(/check)
	for value := range data {
		fmt.Println("Received value", value, ":", time.Now())
	}
	//the reason it got failed is because we loop over another time when we do not have a value
	//for loop is continuously creating new receivers and as soon as it creates new receivers which does not have a sender
	//in this case it'll either block or will result in an error
	//to prevent it we can close the channel for any further values
	//close(data)

	/*
		Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
	*/

}
