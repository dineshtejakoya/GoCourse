package basics

import "fmt"

var middleName = "Cane"

func main() {

	//uninitialized variable
	// var age int
	// the type of the variable is optional if we are initalizing the variable
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10 //":=" shorthanding we can call it as gopher symbol
	// lastName := "Smith"
	//middleName = "William"
	fmt.Println(middleName)
	//Default Types:
	/*
		Numeric Types : 0
		Boolean Types: false
		String Types: ""
		Pointers, slices,maps,functions & Structs: nil
	*/

	//Scope
	//Variables in go have block scope

}

func printname() {
	firstname := "Michael"
	fmt.Println(firstname)
}
