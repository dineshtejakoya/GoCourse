package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//Consistent lock orders help avoid deadlocks
//Go Doesn't have a built in deadlock detection tool like race condition,We can use profiling and debugging tools to identify problematic code patterns
//or use the runtime stack to identify after it got failed

func main() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine1 finished")
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Goroutine2 finished")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function Completed")

	//to check,below statement waits indefinitely until the goroutines are completed
	//select {}

}

//instead of keeping program in a limbo,it'll tells us there is a program with our program
