package main

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

	//Notify channel on interrupt or terminate signals
	//In short subscribe to signals over here using Notify
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGTERM)

	go func() {
		//sig := <-sigs

		for sig := range sigs {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("Received SIGINT (INTERRUPT)")
			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM (TERMINATE)")
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP (Hangup)")
			case syscall.SIGUSR1:
				fmt.Println("Received SIGUSR1 (User defined Signal 1)")
				fmt.Println("User define function is executed")
			}
		}

		// fmt.Println("Received Signal:", sig)
		// fmt.Println("Graceful exit.")
		// os.Exit(0)
	}()

	//Simulate some work
	fmt.Println("Working..")
	for {
		time.Sleep(time.Second)
	}
}

//NOTE:
//TO send interrupt signal we can type ctrl+C in the terminal
//or type kill -s SIGINT <process id>
//to send terminate signal:
//kill -s SIGTERM <processid>
//to send hang signal: kill -s SIGHUP <process-id>
//here as we placed for loop for all signals and handling with switch case
//even though the function supposed to end when the signal is send such as sigint,sigterm,sighub
// we handled them to print or do something else
