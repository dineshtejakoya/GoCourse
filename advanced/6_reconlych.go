package advanced

import "fmt"

func main() {
	//<-chan int  --> receive only channel
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	receiveData(ch)

}

//Receive only channel
//means a channel from which we can only receive data
func receiveData(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
