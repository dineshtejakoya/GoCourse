package intermediate

import "fmt"

func main() {

	// Printing Functions
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print(12.456)

	fmt.Println("Hello")
	fmt.Println("WOrld")
	fmt.Println(12.456)

	name := "John"
	age := 25
	fmt.Printf("Name : %s, Age : %d\n", name, age)
	fmt.Printf("Binary: %b, Age: %X\n", age, age)

	//Formatting Functions
	//useful to concatenate strings
	//it doesn't have new line character at the end
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Print(s)

	//automatically adds spaces
	//it does has new line character at the end
	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	//sprintf formats according to a format specifier and results the string
	//it doesn't have any new line character at the end
	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	//Scanning Functions
	var name1 string
	var age1 int

	fmt.Print("Enter your name and age:")
	//it accepted the first legible(acceptable) value as the first argument
	//fmt.Scan(&name1, &age1)

	//it is similar to scan but it stops scanning at a new line
	//fmt.Scanln(&name1, &age1)

	//it wants to enter our input in the exact format
	fmt.Scanf("%s %d", &name1, &age1)
	fmt.Printf("Name: %s, Age: %d\n", name1, age1)

	//Error Formatting Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
