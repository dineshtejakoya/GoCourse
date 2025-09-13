package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

//Store some values in the context and also use that context with time out

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		//This loop is going to check a value after every 500 milliseconds,fixing the speed of infinity loop using time.sleep
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	//context.Background() is the root context from which other contexts are derived
	//has no deadlock or no timeout
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//instead of deferring cancel we are going to simulate that , we'll perform a heavy task that is going to consume 2 seconds and once that task is finished then i'll cancel this context
	go func() {
		time.Sleep(2 * time.Second) //simulating a heavy task. consider this a heavy time consuming operation
		cancel()                    //manually cancelling only after the task is finished
	}()

	ctx = context.WithValue(ctx, "requestID", "hdfjsldfj23u33")
	//we can add any number of key value pairs to a context
	ctx = context.WithValue(ctx, "name", "John")
	ctx = context.WithValue(ctx, "IP", "11.13.13.55")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	//by the time the ctx reaches here it is cancelled,but it'll retain the values and we can access it from a context which has been cancelled
	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}

	logWithContext(ctx, "This is a test log message")
}

// Contextual Logging
func logWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v", requestIDVal, message)
}
