package basics

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	y := x
	x = append(x, 4)
	fmt.Println(x)
	fmt.Println(y)
	y = append(y, 5)
	fmt.Println(x)
	fmt.Println(y)
	x[0] = 0
	fmt.Println(x)
	fmt.Println(y)
	// fmt.Println("length", len(x), len(y))
	// fmt.Println("Capacity", cap(x), cap(y))

}
