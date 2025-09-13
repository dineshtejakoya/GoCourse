package advanced

//why two struct{}{} showing up while assigning value to channel
import (
	"fmt"
	"time"
)

//channels help goroutines by providing a mechanism to block and unblock goroutines based on channel state

// Channel Synchronization ensures that data is properly exchanged between goroutines
// coordinates the execution flow to avoid race conditions and ensure predictable behavior
// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		//sending value into a channel
// 		done <- struct{}{}
// 	}()

// 	//receiver
// 	<-done
// 	fmt.Println("Finished..")

// }

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Sending..")
		ch <- 9 //Blocking until a value is received
		//here we notice the sent value is not printed,
		// As soon as the value is received by the channel in the main function, the execution is so fast that the program
		// does not leave a time margin fo the goroutine to execute printing the "Sent Value" statement
		time.Sleep(1 * time.Second)
		fmt.Println("Sent Value")
		//to get this printed the main thread(main function) needs to be busy doing something while this statement in go routine executes.
		//You could try putting the main thread to sleep for 1 second before it ends, to simulate some work and see it printed
	}()

	value := <-ch //Blocking until a value is sent
	//time.Sleep(1 * time.Second)
	fmt.Println(value)
}
