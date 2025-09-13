package advanced

import (
	"fmt"
	"time"
)

func main() {
	//Blocking on send only if the buffer is full
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received s:", <-ch)
	}()
	fmt.Println("Blocking starts:")
	ch <- 3 //Blocks because the buffer is full
	fmt.Println("Blocking ends:")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)

}
