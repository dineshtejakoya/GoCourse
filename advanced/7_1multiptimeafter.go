package advanced

import (
	"fmt"
	"time"
)

// We use time.after function to implement time outs providing us a way to handle operations that take too long
func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		//when we are done sending values into a channel close the channel
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}

}
