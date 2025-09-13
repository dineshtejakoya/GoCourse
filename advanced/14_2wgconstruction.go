package advanced

//Try ticketing system using waitgroup
import (
	"fmt"
	"sync"
	"time"
)

// CONSTRUCTION Example
type Worker struct {
	ID   int
	Task string
}

// PerformTasks simulates a worker performing a task
func (w *Worker) PerformTASK(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func main() {
	//create waitgroup
	var wg sync.WaitGroup

	//Define tasks to be performed by workers
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		//wg.Add is not being done inside a goroutine
		//we are adding wg.Add one by one,instead of adding altogether
		wg.Add(1)
		go worker.PerformTASK(&wg)
	}

	//Wait for all workers to finish
	wg.Wait()

	//Construction is finished
	fmt.Println("Construction finished")
}
