package intermediate

import "fmt"

//Each time you call adder(), it returns a new function with its own independent i.
//Closures capture variables by reference (not by value), allowing them to maintain state across invocations.
func main() {
	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	//here we noticed the i value is not resetting to 0
	//Adder function is called only once which is why i:=0 only once
	//we return an anonymous function from adder function that increments i each time its called and returns the updated value
	//whenever we called the returned function sequence,  it increments i and returns the updated value

	//now when i call adder again it will initialize to 0
	//and each time i ran sequence2() function it will run the closure returned from that function
	//in simple closure returned a function here and stored in a variable each time we call that variable
	//  the returned closure is invocoated without affected the adder function
	sequence2 := adder()
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	//anonymous function which is returning another anonymous function which is returning int value
	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	//Using the closure subtracter
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))
	fmt.Println(subtractor(5))
	fmt.Println(subtractor(6))

}

// function name is adder and it is not taking any arguments
// return value is func() with a return value int
func adder() func() int {
	i := 0
	fmt.Println("Previous Value of i:", i)

	//anonymous function
	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}

}
