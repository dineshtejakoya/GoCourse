package basics

import "fmt"

//instead of using if else multiple times in order improvise the readability of the code we can use switch statement
func main() {
	//In switch unlike other lang we don't need to use break it'll be default
	//in case if you want to check the next values also you can use fallthrough

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown Fruit")
	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("Its a weekday")
	case "Sunday":
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Invalid day")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is 20 or more")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

//for no explicitly defining any value we use interfaces
//As per GO compiler we cannot use fallthrough when we are using typeswitch
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("Its float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
