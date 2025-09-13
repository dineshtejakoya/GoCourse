package basics

//import "encoding/base64"
import (
	"errors"
	"fmt"
)

func main() {

	//func functionName(parameter1 type1,parameter2 type2,...) (returnType1,returnType2,..){
	//code block
	//return returnvalue1, returnValue2,...
	//}

	q, r := divide(10, 3)
	//%v is a general placeholder for any values
	fmt.Printf("Quotient :%d. Remainder:%v\n", q, r)

	//Biggest Benefit of multiple return values is error handling they can return an error
	// with values along with values that the function is supposed to return

	result, err := compare(2, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	//named return values
	q1, r1 := divide1(10, 3)
	//%v is a general placeholder for any values
	fmt.Printf("Quotient :%d. Remainder:%v\n", q1, r1)
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// named return values
// in the return values we are mentioning quotient and remainder
// also we don't need to explicitly mention return quotient, remainder, go is smart enough to understand that
// its gonna return both if we just place return as we have explicitly mentioned them in the function
func divide1(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		//in the function declaration we have explicitly mentioned that return type would be string and error
		//if we don't have any error to send we can send it as nil
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		//here we are trying to send only error
		//return nil,
		//here we cannot send nil because error is an interface we can use nil i.e., the zero value for interface is nil
		//similarly the zero value for string is a blank string
		return "", errors.New("Unable to compare which is greater")
	}
}
