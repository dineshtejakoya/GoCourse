package intermediate

import "fmt"

//embedding
type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with Value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

//Method with Pointer Receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor //shorthand
	r.width *= factor
}

func main() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle width width 9 and length 10 is", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle with a factor of 2", area)

	num := MyInt(-5)
	num1 := MyInt(5)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	//if we were to call area it would be s.rectange.area
	//but now because we are embedding a struct inside another struct
	// the method associated with the embedded struct will be promoted to the outer struct
	fmt.Println(s.Area())
	//both are same
	fmt.Println(s.Rectangle.Area())
}

//Methods can be associated with any type, it's not just for structs
type MyInt int

//Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

//this time we didn't create an instance
//it's not a thumb rule to use instance
//here we aren't accessing any value
func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}
