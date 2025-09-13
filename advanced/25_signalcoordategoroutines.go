package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//IGNORE THE SQUIGGLE LINE UNDER SIGUSR1, IT WORKS AS EXPECTED

func main() {
	//to get the process id
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	//create signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//Notify channel on interrupt or terminate signals
	//In short subscribe to signals over here using Notify
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		fmt.Println("Received Signal:", sig)
		done <- true
	}()

	go func() {
		//sig := <-sigs
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}

	}()

	//Simulate some work
	//fmt.Println("Working..")
	for {
		time.Sleep(time.Second)
	}
}
