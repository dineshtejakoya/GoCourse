package basics

import (
	"fmt"
	"maps"
)

// Maps are built in data structures that associate keys with values
// These are like dictionaries in other programming languages, and provide an efficient way to lookup data by a key
// Maps provide an efficient way to store and retrieve key value pairs
// Each key must be unique with in the map
// And keys are typically of comparable type like strings,integers
// Maps are unordered collection of key value pairs meaning there is no guaranteed order while iterating over them
func main() {
	//var mapVariable map[keyType]valueType
	//we can also use make to declare a map
	//mapVariable = make(map[keyType]valueType)
	//using a map Literal
	// mapVaraiable = map[keyType]valueType{
	// 	key1:value1,
	// 	key2:value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	//To retrieve value based on key
	fmt.Println(myMap["key1"])
	//If the Key doesn't exist it'll return a 0 value
	fmt.Println(myMap["key"])
	//we can change the values for a key
	myMap["code"] = 21
	fmt.Println(myMap)

	//If we want to delete a key value pair we can use delete function
	delete(myMap, "key1")
	fmt.Println("New Map after deleting", myMap)

	myMap["key2"] = 24
	myMap["key3"] = 26
	myMap["key4"] = 20
	fmt.Println("New Map", myMap)

	//If we are not looking for the value for a specific key, we just want to check, we just want to check if there is any value associated with the key
	//we can try and get the second value which is an optional usable value
	value, ok := myMap["key2"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	fmt.Println(value)
	fmt.Println("Is a value associated with key1", ok)

	//To delete all the key value pairs we use clear function over maps
	clear(myMap)
	fmt.Println("Map After deleting all key value pairs", myMap)

	//New map declaration
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	//To iterate over the map we use for loop with range
	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	//Suppose we have a map that hasn't bee initialized but declared then it is nil
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil")
	}

	val := myMap4["key"]
	fmt.Println(val)

	/*Below won't allow you to create key value pairs for myMap4 as we have initialized the map with nil,
	instead we need to use make to do that, its something unique with go
	myMap4["key"] = "Value"
	fmt.Println(myMap4)
	*/

	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	fmt.Println("myMap3 length is", len(myMap3))

	//Nested Maps
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
