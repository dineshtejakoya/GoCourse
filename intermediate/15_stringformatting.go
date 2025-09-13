package intermediate

import "fmt"

func main() {
	num := 42
	fmt.Printf("%05d\n", num)

	message := "Hello"
	//we are using pipe symbol just to see start and end
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	//go supports string interpolation using ``(backticks)
	//backtick makes a string a raw, not allowing escape sequences to execute
	message1 := `Hello \n World`
	message2 := "Hello \n World"
	fmt.Println(message1)
	fmt.Println(message2)
}
