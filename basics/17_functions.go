package basics

import "fmt"

//encapsulating reusable code blocks
func main() {
	//func <name>(parameters list) returnType{
	//code block
	//return value
	//}

	//When we are making public function it should start with Upper Case Letter
	//Incase of private function we need to start with lower case letter
	//ex:fmt.Println() it allows Println() function to be used by other packages to be used this function

	//return statement is optional and it returns one or more values to the caller
	//if return statement is omitted, functions returns default zero values for their types

	//sum := add(1, 2)
	fmt.Println(add(2, 3))

	//Anonymous Functions, Closures or Function Literals
	func() {
		fmt.Println("Hello Anonymous Function")
	}()
	//() is called immediately

	greet := func() {
		fmt.Println("Assigning Function to Greet Variable")
	}

	greet()

	//we can use functions as types
	//Functions in go can be assigned to variables, passed as arguments to other functions and returned from functions
	operation := add

	result := operation(3, 5)
	fmt.Println(result)

	//Passing a functionas argument
	result1 := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result1)

	//Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b

}

//Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

//Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
