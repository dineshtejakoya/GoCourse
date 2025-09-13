package advanced

import (
	"fmt"
	"time"
)

func main() {
	// //========Blocking on Receive only if the buffer is empty
	// ch := make(chan int, 2)
	// //blank channel print cause error

	// //channels they wait for goroutine to finish
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- 1
	// 	ch <- 2
	// }()
	// fmt.Println("Value: ", <-ch)
	// fmt.Println("Value: ", <-ch)
	// fmt.Println("End of Program")

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	//when this goroutine is extracted by the go runtime scheduler
	//it extracts from the main thread
	go func() {
		//fmt.Println("Blocking function")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
		//we noticed sometimes the value "Received: 1" is not printed in the console
		//it's because "fmt.Println("Received:", <-ch)" execution starts from end <- starts
		//in the meantime <-ch is executed and the blocking session ch -> 3 unblocks and executes
		/*Go kills all goroutines when the main goroutine exits.
		There’s no guarantee your goroutine will run after main finishes unless you explicitly synchronize*/
	}()

	fmt.Println("Blocking starts")
	ch <- 3
	fmt.Println("Blocking ends")

}
