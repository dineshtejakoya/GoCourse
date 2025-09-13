package basics

import "fmt"

//The range keyword in go provides a convenient way to iterate over data structures like arrays,slices,strings,maps and channels
//iterates and access elements
func main() {
	//to iterate range over strings
	//we see unicode characters in the below case for value v
	message := "Hello World"
	for i, v := range message {
		fmt.Println(i, v)
	}

	for i, v := range message {
		//The characters/Alphabets are called runes in go
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	//Range keyword operates on a copy of the data structure it iterates over,
	// therefore modifying index or value inside the loop does not affect the original data structure

	//For Arrays, Slices and Strings range iterates in order from the first element to the last
	//For Maps range  iterates over key value pairs, but in an unspecified order
	//For channels range iterates until the channel is closed

	//Using blank identifier such as underscore(_), this practice prevents memory leaks by allowing Go's garbage collector to reclaim the memory
	//It is necessary if we are not using any values that we are receiving from a map or array slice
	// anything then we use blank identifier such as underscore to prevent memory leaks
}
