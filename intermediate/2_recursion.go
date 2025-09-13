package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
	fmt.Println(sumofDigits(256))
	fmt.Println(sumofDigits(123456789))

}

func factorial(n int) int {
	//Base Case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive Case: factorial of n is n * factorial(n-1)
	return n * factorial(n-1)
}

func sumofDigits(n int) int {
	//Base Case
	if n < 10 {
		return n
	}

	return n%10 + sumofDigits(n/10)
}

//sometimes recursion techniques can be optimized using memoization
//caching results of expensive function calls
