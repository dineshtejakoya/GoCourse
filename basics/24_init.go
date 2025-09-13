package basics

import "fmt"

//init is a special function that can be declared in any package
//its used to perform initialization tasks for the package before it is used
//go executes initi automatically when the package is initialized
//The init function always gets executed before main function
// it occurs exactly once per package even if the package is imported multiple times

func init() {
	fmt.Println("Initializing Package1...")
}

func init() {
	fmt.Println("Initializing Package2...")
}

func init() {
	fmt.Println("Initializing Package3...")
}

func main() {

	fmt.Println("Inside the main function")
}
