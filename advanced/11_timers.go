package advanced

import (
	"fmt"
	"time"
)

/*
A timer in GO allows you to schedule an event to occur after a specified duration
It is useful for implementing timeouts, scheduling periodic tasks and delaying operations
Key reason:
Timeout-> to limit how long a particular operation should wait
Delays -> Schedule operations to occur after a certain delay, we have been using time.sleep but we can use timers instead as well
Periodic Tasks -> execute recurringly at regular intervals

We create a timer using time package ->new Timer
*/
func main() {
	fmt.Println("Starting app..")
	//NewTimer is non blocking in nature, time.sleep() is blocking in nature
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for timer.C")

	//to stop the timer before it expires
	stopped := timer.Stop()

	if stopped {
		fmt.Println("Timer stopped")
	}

	//it'll restart a timer that has been expired or stopped
	timer.Reset(time.Second)
	fmt.Println("Timer reset")
	//here the C field of the timer type is a channel that receives the time when the timer expires
	<-timer.C //blocking in nature
	fmt.Println("Timer expired")

}
