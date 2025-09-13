package advanced

import (
	"context"
	"fmt"
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
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second) //timer of context starts here.You have this specified time duration to use this context.
	//After this time the context will send a cancellation signal
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "hdfjsldfj23u33")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	//by the time the ctx reaches here it is cancelled,but it'll retain the values and we can access it from a context which has been cancelled
	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}
}
