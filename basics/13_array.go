package basics

import "fmt"

//An array is fixed size collection of elements of the same size
//allows us to store multiple values under a single variable
func main() {
	//var arrayName [size]elementType
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//standard i,v
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	//if we want to discard the index we can use underscore(_) its a blank identifier
	//because if we use some other variable and not using it , then it will show as error like defined but not used
	//underscore is blank identifier, used to store unused values
	for _, v := range numbers {
		fmt.Printf("Value is %d\n", v)
	}
	a, _ := someFunction()
	fmt.Println(a)

	//Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Array 1 is equal to Array2:", array1 == array2)

	//Go supports multidimensional arrays which are arrays of arrays,
	// they are usef for representing other structured data
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	//using pointers to update arrays in original array
	originalArray1 := [3]int{1, 2, 3}
	var copiedArray1 *[3]int

	copiedArray1 = &originalArray1
	copiedArray1[0] = 100

	fmt.Println("Original Array:", originalArray1)
	fmt.Println("Copied Array:", *copiedArray1)

}

func someFunction() (int, int) {
	return 1, 2
}
