package advanced

import (
	"fmt"
	"time"
)

// =============TIMEOUT
func longRunningOperation() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
func main() {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperation()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Operation times out")
	case <-done:
		fmt.Println("Operation completed")
	}
}
