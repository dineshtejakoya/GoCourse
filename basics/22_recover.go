package basics

import "fmt"

//recover is a built in function that is used to regain control of a panicking goroutine
//It's only useful inside defer functions and used to manage behaviour of a panicking go routine to avoid abrupt termination
func main() {

	//Panic is a built in function that stops the ordinary flow of control and begins panicking
	//When the function panic is called , the current function stops execution,
	//  and any deferred functions in that function are executed and then the control returns to the calling function

	//Recover is to regain control of the panicking go routine
	//Its a built in function that stops the propagation of the panic and returns the value passed to the panic call
	process()
	fmt.Println("Returned from Process")

}

func process() {
	defer func() {
		//r := recover()
		//if r!=nil{}
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}
