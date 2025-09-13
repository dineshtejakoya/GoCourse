package advanced

import (
	"fmt"
	"sync"
	"time"
)

// EXAMPLE with channels
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	//first this defer wg.Done()
	defer wg.Done()
	fmt.Printf("WorkerID %d starting. \n", id)
	time.Sleep(time.Second) //simulate some work
	for task := range tasks {
		results <- task * 2
	}

	fmt.Printf("WorkerID %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	//Add wg.Add
	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i+1, tasks, results, &wg)
	}

	//passing values to tasks channel
	for i := range numJobs {
		tasks <- i + 1
	}

	//Now we need to close the tasks channel as well
	//we need to do it once we have sent all that values to that channel
	close(tasks)

	//As results is a channel, we need to close it once it is done
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
