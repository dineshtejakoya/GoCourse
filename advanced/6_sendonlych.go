package advanced

//This example is send only channel
import "fmt"

//Channel Directions specify the allowable operations on a channel either sending or receiving
//Channel Directions are intended for use in functions and goroutines

/*
Why Channel Directions?
Improve code quality and maintainability
Prevent unintended operations on channels
Enhance type safety by clearly defining the channels purpose
*/

/*
By default a channel is bi-directional
We can turn a channel into uni-directional
*/
func main() {
	//<-chan int  --> receive only channel
	ch := make(chan int)

	//go routine which accepts a channel which is send only channel
	//here we are restricting the use of our channel inside our goroutine or function
	go func(ch chan<- int) {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}(ch)

	for value := range ch {
		fmt.Println("Received: ", value)
	}

}
