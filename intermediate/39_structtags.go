package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

//Tags play a crucial role how data is encoded and decoded in go, especially when working with json
//struct tags can be used to map struct field names to specific json keys, which might not match the go field names
//useful when working with APIs , datasources where the json keys have different naming conventions
//struct tags also indicate that certain fields should be omitted from the json o/p either when they have 0 values or always
//for omitting zero values we use omitEmpty
//for always omitting the field we use -

type Person struct {
	//tags for json,db and xml
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"-"`
}

//to permanently omit a value we use -

func main() {
	//person := Person{FirstName: "Jane", LastName: "", Age: 50}
	person := Person{FirstName: "Jane", Age: 30}
	jsonData, err := json.Marshal(person)

	if err != nil {
		log.Fatalln("Error marshalling struct", err)
	}
	fmt.Println(string(jsonData))
}
