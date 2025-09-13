package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	/*PascalCase ==> it is typically used for typing names in go
	such as structs, interfaces, enums
	it distinguishes typenames from variables and functions
	it helps maintain clarity and consistency in the code
	Eg: CalculateArea, UserInfo, NewHttpRequest
	*/

	/* snake_case
	it is commonly used for naming variables, constants and filenames
	eg: user_id, first_name, http_request
	*/

	/*UPPERCASE
	it is used exclusively for naming constants in go
	to emphasize their immutability
	eg: ROLLNUMBER
	*/

	/*mixedCase
	Eg: javaScript, htmlDocument, isValid
	*/

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}
