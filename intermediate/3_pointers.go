package main

import "fmt"

//A pointer is a variable that stores the memory address of another variable
//Note: Go doesn't support pointer arithmetic
//pointer operations are limited to referencing and dereferencing in go
func main() {

	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(ptr)
	fmt.Println(*ptr) //dereferencing a pointer

	//A pointer that doesnot point to any memory address is a nil pointer

	// var a int = 10
	// fmt.Println(a)
	// fmt.Println(&a)

	// var b int = 10
	// fmt.Println(b)
	// fmt.Println(&b)

	var ptr1 *int
	var b int = 10
	fmt.Println(b)
	if ptr1 == nil {
		fmt.Println("Pointer is nil")
	}

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
