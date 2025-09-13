package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go!"

	fmt.Println(len(str))
	//concatenation
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	//get index of any letter any rune
	fmt.Println(str[0]) //get ascii value

	//slicing
	fmt.Println(str[1:])
	fmt.Println(str[1:7])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	//we know that length is only for array or a slice, not for a number
	fmt.Println(len(str3))

	//strings splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ", ")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	//concatenates elements of a slice into a single string with a separator
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	//contains
	fmt.Println(strings.Contains(str, "Go"))

	//replace
	//ending contains number of occurences we want to replace
	replaced := strings.Replace(str, "Go", "Dinesh", 1)
	fmt.Println(replaced)

	//trim leading and trailing spaces
	strwspace := " Hello Everyone! "
	fmt.Println(strings.TrimSpace(strwspace))

	//Upper to lower and lower to upper
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	//repeat function
	fmt.Println(strings.Repeat("foo ", 3))

	//we can also count the occurence of an alphabet or a substring inside another string
	fmt.Println(strings.Count("Hello", "He"))

	//check prefix or suffix
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	//Regular Expressions
	str5 := "Hello, 123 456 Go"
	//patterns need to be declared inside backticks(``)
	re := regexp.MustCompile(`\d+`)
	//below ending value is for number of matches, if we want all we use -1
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)
	for _, v := range matches {
		fmt.Println(v)
	}

	//we have another package that lets us work on unicode characters and strings
	//i.e., unicode utf 8 package
	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6))

	//Strings Builder, alternative as strings are usually immutable
	//more efficient, it minimizes memory allocation and copies
	//builder can be used immediately after declaration without initialization
	//can be reused by calling its reset method which clears its internal buffer
	var builder strings.Builder
	//Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//Convert builder to a string
	result2 := builder.String()
	fmt.Println(result2)

	//using Writerune to add a character
	//whenever we are dealing with runes it has to be single quotes
	builder.WriteRune(' ')
	builder.WriteString("How are you")
	result2 = builder.String()
	fmt.Println(result2)

	//Reset the builder
	builder.Reset()
	builder.WriteString("Starting Fresh")
	result2 = builder.String()
	fmt.Println(result2)
}
