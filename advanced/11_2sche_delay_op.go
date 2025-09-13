package advanced

import (
	"fmt"
	"time"
)

// ========Multiple Timers
func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("TImer1 expired")
	case <-timer2.C:
		fmt.Println("Timer2 expired")

	}
}

// func main() {
// 	timer := time.NewTimer(2 * time.Second) //non blocking timer starts

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed Operation Executed")
// 	}()

// 	fmt.Println("Waiting")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("End of the program")
// }
