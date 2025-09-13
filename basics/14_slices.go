package basics

import (
	"fmt"
	"slices"
)

// Slices are dynamic, flexible views into arrays
// They provide more powerful and convenient interfaces into sequences of data compared to arrays
func main() {
	//slices do not have fixed length
	//var sliceName[]ElementType
	// var numbers []int
	// var numbers1 = []int{1,2,3}

	// numbers2 := []int{9,8,7}
	//Slices are references to underlying arrays
	//They do not store any data themselves but provide a window into the arrays elements
	//Slices can grow and shrink dynamically
	//we have len function to check length of the slice
	//we have cap function to check capacity of the slice
	//we can also initialize slice using make
	//below we are initializing slice with capacity of 5
	// slice := make([]int,5)

	//lets say we have array of 5 elements we can make slice out of the existing array
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("Slice1", slice1)

	//copy slice
	//to copy we use make function
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Copied Slice", sliceCopy)

	//NilSlices has a zero value and doesnot reference any underlying array
	//var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice1", slice1[3])

	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slice1 is equal to sliceCopy")
	}

	//multidimesnsional slice(2d slice)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	fmt.Println("slice1:", slice1)

	slice2 := slice1[2:4]
	fmt.Println(slice2)

	//capacity of the slice is sum of the length of the slice and length of the array beyond the slice
	fmt.Println("The capacity of slice1 is", cap(slice1))
	fmt.Println("The length of slice2 is", len(slice2))
}
