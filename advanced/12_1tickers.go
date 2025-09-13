package advanced

import (
	"fmt"
	"time"
)

//==============SCHEDULING Logging, Periodic Tasks, Polling for updates
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// Polling for updates
// if we want to stop the ticker and print a message and perform the action once the ticker stops
func main() {
	ticker := time.NewTicker(time.Second)

	//time.After() returns a channel once expires
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		//stop from time.After
		case <-stop:
			fmt.Println("Stopping ticker..")
			return
		}
	}

}
