package advanced

import (
	"fmt"
	"time"
)

//A worker pool is a design pattern used to manage a group of workers(i.e., goroutines)
//this pattern helps in handling number of concurrent goroutines and controlling tasks
//used for resource management,limits the number of concurrent goroutines to avoid overwhelming system resources,limit the number of concurrent goroutines according to our will
//using workerpool we efficiently distribute the tasks among a fixed number of workers
//we can now scale the processing of tasks without creating an excessive number of goroutines

/*
Building Blocks
Tasks
Workers
Task Queue


Implementation Steps:
create task channel
create worker goroutines
distribute tasks
graceful shutdown
*/

// ==========BASIC Worker Pool Pattern
// worker function
// channel from which we receive the tasks
// another channel to which we send the processed tasks
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		//Simulate work
		time.Sleep(time.Second)
		results <- task * 2
	}
}
func main() {
	//fix the number of workers
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//Create workers
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	//we need to send values to the tasks channel for the workers to process those tasks
	//Send values to the tasks channel
	for i := range numJobs {
		tasks <- i
	}

	close(tasks)

	//collect the results
	for range numJobs {
		result := <-results
		fmt.Println("Result:", result)
	}
}
