package advanced

import (
	"fmt"
	"time"
)

// ===========Synchronizing Multiple Goroutines and Ensuring all goroutines are complete
func main() {
	//Create a buffered channel
	numGoroutines := 3
	done := make(chan int, 3)

	//we are going to use a loop
	//we'll create goroutines inside the loop
	for i := range numGoroutines {
		//this function is going to take a goroutine id
		time.Sleep(2 * time.Second)
		go func(id int) {
			fmt.Printf("Goroutine %d working ..\n", id)
			//time.Sleep(time.Second)
			//once our goroutine is finished, we send a signal to our channel that it is finished
			done <- 1
		}(i)
	}

	//2 go routines will get created and third one won't be blocking because <-done is not there to block third time to make sure the goroutine to get finished
	//
	// for range 2{
	// 	<-done
	// }

	for range numGoroutines {
		//wait for each goroutine to finish
		<-done
	}
	fmt.Println("All goroutines are finished:")

}
