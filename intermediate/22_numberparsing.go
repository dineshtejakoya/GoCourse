package intermediate

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	//Atoi means Alphabets to integers, of type int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed Integer:", num+1)

	//parseint , it converts a string to an integer with specified base and bitsize
	//suppose i want integer to be of type int32 or int64
	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed Integer:", numistr)

	//parseFloat
	floastr := "3.14"
	floatval, err := strconv.ParseFloat(floastr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return
	}
	fmt.Printf("Parsed Float: %.2f\n", floatval)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}

	fmt.Println("Parsed binary to decimal:", decimal)

	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing hexadecimal value:", err)
		return
	}

	fmt.Println("Parsed hex to decimal:", hex)
}
