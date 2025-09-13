package advanced

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come
// back to join the main thread once the functions are finished/ready to return any value
// Goroutines do not stop the program flow and are non blocking
// Goroutine exits when the function completes, go run time manages go routine scheduling and execution
// go uses m:n scheduling model, where m goroutines are mapped to n operating system threads
// multiplexing is kinda like switching, go routine scheduler multiplexes or switches go routines onto available os threads
func main() {
	var err error
	//our program ended before the goroutine could comeback to the main thread
	//program cease to exist by the time sayhello() finished
	//in this case we need to wait for goroutine to finish
	//basic method to wait for goroutine to finish is time.sleep()
	fmt.Println("Beginning Program:")
	go sayHello()
	fmt.Println("After sayhello function")

	// err = doWork() //this is not accepted
	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

// Basic goroutine in action
func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		//fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

//Concurrency means multiple tasks progress simultaneously and not necessarily at the same time
//but parallelism states that tasks are executed literally at same time on multiple processors
//so goroutines facilitates concurrency and goruntime schedules them across available cpus for parallelism when possible
//how goroutines handle error back to the main thread: use return values or shared error variables if not using channels

func doWork() error {
	//Simulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in dowork")
}

//when using goroutines in application, try and limit the creation of goroutine,
//excessive creation of goroutine cann lead to resource exhaustion
