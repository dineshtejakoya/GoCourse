package advanced

import (
	"fmt"
	"sync"
	"time"
)

//WaitGroup is a synchronization primitive provided by sync package in Go.
//It is used to wait for collection of goroutines to complete their execution
//(We have another mechanism to to wait for goroutines to finsh other than using channels)
//we use it for synchronization
//wait for multiple goroutines to finish before proceeding
//Also help us in coordination

/*
To create instance of wait group we use sync package and the function is sync.waitGroup

Methods:
Add(delta int) increments waitgroup counter, typically it is used to indicate the number of goroutines to waitfor
Done() decrements the counter by one,should be called inside goroutine when it finishes its task
Wait() method blocks until the counter inside the waitgroup is decremented to zero

*/

// =========BASIC Example withou Using Channel
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	//wg.Add(1) wrong practice
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) //Simulate some time spent on processing the task
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers) //This is the correct way of adding counter to wait groups

	//Launch workers
	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait() //blocking mechanism until wait group decrements
	fmt.Println("All workers finished")
}
