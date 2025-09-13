package basics

import (
	"fmt"
	"math"
)

func main() {
	//Variable Declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	//in order to get floating point value through division(/)
	// you need to make sure atleast one value is in fractional part
	//the reason is it starts from right to left, it didn't see initially that the p value is float
	const p float64 = 22 / 7.0
	fmt.Println(p)

	//Overflow occurs when the result of a computation exceeds the maximum value that can be stored in a given numeric data type
	//In GO Overflow results in the value wrapping around to the minimum value around signed integers or causing unexpected behaviour for unsigned integers
	//underflow occurs when the result of a computation is smaller than the minimum value that can be stored in a given numeric data type, this is more relevant for floating point numbers where underflow can lead to loss of precision or significant digits in calculations involving very small values
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	//Overflow with Signed Integer
	maxInt = maxInt + 1
	//it will result in negative value
	fmt.Println(maxInt)

	//Overflow with Unsigned integer
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	//Underflow example
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
