package basics

import "fmt"

//In Go panic is a built in function that stops normal execution of a function immediately,
// when a function encounters panic, it stops executing
//its current activities, unwinds the stack, and then executes any deferred functions
func main() {
	//A panic is typically used to signal an unexpected error condition where the program cannot proceed safely
	//syntax: panic(interface{})
	//interface means you can input any value of any type, as an argument for this function

	//Example of valid input
	process(10)

	//Example of invalid input
	process(-3)

	//lets see defer with pani, defer will execute when the function returns the value
	//but it'll also execute even when the function is panicking

}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("Input must be a non negative number")
		//below is an unreachable code because as soon as we encounter panic anything after that will not be executed
		//fmt.Println("After Panic")

		//defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing Input:", input)
}
