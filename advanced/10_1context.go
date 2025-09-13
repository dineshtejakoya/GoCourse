package advanced

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int) string {
	//with contexts we have cancellational signal operate like below
	select {
	//ctx.Done() is function which sends completion signal or cancellation signal
	case <-ctx.Done():
		return fmt.Sprintf("operation cancelled: %v", ctx.Err())
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result from timeout context:", result)

	//lets make our function sleep for 2 seconds
	//as the timeout has surpassed 1 second , the checkEvenOdd function will throw error
	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result after timeout", result)
}
