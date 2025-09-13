package advanced

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		//Closing a channel signals that no more values will be sent on that channel
		close(ch)
	}()

	//if we comment out the if condition it'll run in infinity loop, because the channel is closed but we are still looping over
	for {
		select {
		//ok  is a boolean value here to indicate if a channel is open or closed
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel is Closed")
				//clean up activities
				return
			}
			fmt.Println("Received:", msg)
		}
	}

}
