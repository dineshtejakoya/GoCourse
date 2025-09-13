package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

// marshall converts Go data structures into JSON
// unmarshall converts json into go data structures(decoding)
// commonly used for transmitting data in web applications
// encoding/json ->encode(Marshal) ,decode(deMarshal)
type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

//omitempty means the field will be excluded from the JSON if it has the zero value for its type
// (0 for ints, "" for strings, nil for pointers/slices/maps, etc.).

// backticks(`) in structs are primarily used for struct field tags which provides meta data about the fields
// especially useful when you need to conver the struct to json or interact with databases
// in json fields are lower that's why in backticks name and age are in lowercase
// as both field are string we are using double quotes
// for database orm we use `db:"user_id"`
func main() {
	person := Person{FirstName: "John", Age: 35}

	//Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakemail.com", Address: Address{City: "New York", State: "NY"}}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name":"Jenny Doe","emp_id":"0009","age":30,"address":{"city":"San Jose","state":"CA"}}`

	var employeeFromJson Employee
	//Decode the json string
	//first variable is data, second variable is the data that it is going to store
	//and also it should needs to be a pointer
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(employeeFromJson)
	//we can access the specific values of the struct
	fmt.Println("Jenny's age increased by 5 years", employeeFromJson.Age+5)
	fmt.Println("Jenny's City", employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "san Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error marshalling json", err)
	}

	//here we receive slice of byte(numbers-see)
	//fmt.Println(jsonList)
	//those slices we need to convert to string
	fmt.Println(string(jsonList))
	//then we can notice after marshalling field names will come between {}

	//Handling unknown JSON structures
	jsonData3 := `{"name":"John","age":30,"address":{"city":"New York","state":"NY"}}`
	//because it is unknown json structure we are going to make a map

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
