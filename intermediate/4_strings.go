package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	//\r is carriage it takes the cursor to the first position in line
	message2 := "Hello, hi\rGo!"
	rawMessage := `	Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	//strings are made up of rune characters
	fmt.Println("Length of message variable is", len(message1))
	fmt.Println("The first character in message var is", message[0]) //ASCII Character of H

	//concatenation
	greeting := "Hello "
	name := "Alice"
	msg := greeting + name
	fmt.Println(msg)

	//Comparision
	str1 := "Apple"
	str2 := "banana"
	fmt.Println(str1 < str2) //lexicographical comparision(ascii value)

	for _, char := range message {
		//%d,%c are formatting verbs or placeholders
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%x\n", char)
		//%v gives ascii value in uint8 format(ascii value)
	}

	//To count number of runes in a string
	fmt.Println("Rune Count:", utf8.RuneCountInString(greeting))

	//we have to manipulate strings by creating new strings, because strings are immutable
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	//in go rune is an alias for int32
	//runes are used to represent individual character in a string
	var ch rune = 'a'
	jch := 'D'

	//this will print ascii character
	fmt.Println(ch)
	fmt.Println(jch)

	//We will get actual characters here
	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	//Convert Runes to string
	cstr := string(ch)
	fmt.Println(cstr)

	//to check if cstr is actually a string use formatting verb
	fmt.Printf("Type of cstr is %T\n", cstr)
	//type of rune ch is int32 because we declared it as rune
	fmt.Printf("Type of cstr is %T\n", ch)

	//Runes facilitate handling of unicode characters supporting internationalization
	//we can use text in any language
	const nihongo = "この服は寒い冬の日には向かない"
	fmt.Println(nihongo)

	jhello := "この服は寒い冬の日には向かない"
	for _, runeValue := range jhello {
		//%v is the default value of the rune values ,which are integer values
		fmt.Printf("%v\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)

	//explore difference between string and runes
}
