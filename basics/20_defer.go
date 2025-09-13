package basics

import "fmt"

//Defer is a mechanism that allows you to postpone the execution of a function until the surrounding function returns
//Its a useful feature ensuring that certain cleanup actions or finalizing tasks are performed regardless of how the function exists,
//  whether it returns normally or panic
func main() {
	//A defer statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the
	// surrounding function executed a return statement reached the end of its statements
	//reached the end of its return body
	//or because the corresponding go routine is panicking

	//Go Routines are functions which run concurrently in the background and they are not part of the main thread
	//So any function which is a go routine is thrown to the back so that it finishes off its work, not blocking the main thread
	//but in the background, and then it coomes back and joins the main thread

	//Syntax: Defer Statement is followed by function or method call
	//Surrounding function means the function that encloses the defer function

	process()
	process1()
	fmt.Println("===========Process2=============")
	process2(10)
}

func process() {
	defer fmt.Println("Deferred statement executed")
	fmt.Println("Normal execution statement")
}

//We can also have multiple deferred statements in a function and they will be executed in last in first out order
func process1() {
	defer fmt.Println("FirstDeferred statement executed")
	defer fmt.Println("Second Deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")
	fmt.Println("Normal execution statement")
}

func process2(i int) {
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("FirstDeferred statement executed")
	defer fmt.Println("Second Deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i", i)
}

//Deferred statement can be compared with finally block in other programming langauges
