package advanced

import (
	"fmt"
	"reflect"
)

// ==========WORKING WITH METHODS ============
type Greeter struct{}

func (g Greeter) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func main() {
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var method reflect.Method

	fmt.Println("Type:", t)
	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}

	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Doe")})
	//[] string{"Alice"}
	//[]type{type("dfdf"), type {"jdhfgkj"}}
	fmt.Println("Greet Result:", results[0].String())
}

//===============METHODS WITH REFLECT -2=============
// type person struct {
// 	//If we are using uppercase of the field then it can be exported, and can be accessed outside the pacakge
// 	Name string
// 	//lowercase age constricts the visibility of the struct and the fields of the struct to it to the package only
// 	age  int
// }

// func main() {
// 	//create instance of person
// 	p := person{Name: "Dinesh", age: 26}
// 	v := reflect.ValueOf(p)

// 	//v.NumField returns the total number fields of instance p
// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}

// 	v1 := reflect.ValueOf(&p).Elem()

// 	//we can change the values of struct as well
// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 	} else {
// 		fmt.Println("Cannot set")
// 	}

// 	fmt.Println("Modified Person:", p)

// }

// ==============VALUES OF INTERFACE TYPE -1
// func main() {
// 	var itf interface{} = "Hello"
// 	v3 := reflect.ValueOf(itf)

// 	fmt.Println("V3 Type:", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String Value:", v3.String())
// 	}
// }

//==========INITIAL REFLECT EXAMPLE -0
// func main() {

// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())
// 	fmt.Println("Is Int:", t.Kind() == reflect.Int)
// 	fmt.Println("Is String:", v.IsZero())

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
// 	fmt.Println("V2 Type:", v2.Type())

//NOTE: it'll panic if it is not an integer.(check)
// 	fmt.Println("Original Value:", v1.Int())

// 	v1.SetInt(18)
// 	fmt.Println("Modified Value:", v1.Int())

// }
