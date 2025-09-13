package basics

import "fmt"

func main() {
	//Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		//Placeholder between double quotes, %v is general value and %d specific to numbers
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break
		}
	}

	rows := 5

	//Pyramid Pattern with Stars
	for i := 1; i <= rows; i++ {
		//Inner Loop
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println() //Move to the next lines
	}

	//enhancement in go 1.22 ability to range over integers
	for i := range 10 {
		fmt.Println("Range", 10-i)
	}
	//Go has no while loop
}
