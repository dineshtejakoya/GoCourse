package basics

import "fmt"

//Variadic functions in go allow you to create functions that can accept a variable number of arguments
//This flexibility is particularly useful when you want to design functions that can handle a varying number of inputs, without specifying them individually
//In go, variadic functions are defined by prefixing the type of the last parameter with an ellipsis(...)
func main() {
	// ... Ellipsis --> this indicates that the function can accept zero or more arguments of that type
	//we have to use ellipsis before declaring the type
	//func functionName(param1 type1,param2 type2,param3 ...type3) returnType{
	//function body
	//}
	//so when we are declaring the parameters of our function with three dots ellipsis, it's called variadic parameter

	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	statement, total := sum1("The Sum of 1,2,3 is", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}
	//... after numbers variable is a variadic expansion operator
	//... after a slice is used to expand/unpack it when calling a variadic function
	sequence3, total3 := sum2(3, numbers...)
	fmt.Println("Sequence: ", sequence3, "Total", total3)

}

//Variadic Parameters must be last parameter in the function signature
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

func sum1(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}

func sum2(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
