package main

import "fmt"

//Struct embedding allows a struct to inherit fields and methods from another struct type
//powerful for code reuse and structuring data
type person struct {
	name string
	age  int
}

//Here employee struct embeds person
//inherits the fields and methods of functions
type Employee struct {
	//employeeInfo person //embedded struct Named field
	person //Anonymous field
	empId  string
	salary float64
}

//Method Inheritance
func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old \n", p.name, p.age)
}

//methods can be overriden by redefining them in the outer struct
func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID %s, and I ear %.2f.\n", e.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		person: person{name: "Teja", age: 25},
		empId:  "228",
		salary: 140000,
	}

	//field promotion will only happen when it is an anonymous struct
	//fmt.Println("Name: ", emp.employeeInfo.age)
	fmt.Println("Name: ", emp.name)
	fmt.Println("Age: ", emp.age)
	fmt.Println("Emp ID: ", emp.empId)
	fmt.Println("Salary: ", emp.salary)

	//introduce is a method of person struct and not of employee
	//but because person is embedded into employee we can directly access introduce from an instance of an employee
	emp.introduce()
}
