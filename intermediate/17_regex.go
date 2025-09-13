package intermediate

import (
	"fmt"
	"regexp"
)

// Regular expressions provide a powerful way to search manipulate and validate text strings based on patterns
// Widely used in text processing, searching and data extraction tasks
// To work with regex pattern in go, you first compile it using the regexp.compile or regexp.Mustcompile
func main() {
	fmt.Println("He said,\"I am great\"")
	fmt.Println(`He said,"I am great"`)

	//compile a regex pattern to match email
	// address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	//Test some strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	//Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	//Capturing Groups
	//Compile a regex to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	//declare test strings
	date := "2024-07-30"

	//Find all submatches
	//The result submatches variable is a slice
	//the first element at index 0 is the source
	//then starting from index 1 are the submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	//Source String
	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	//Flags and options
	//i - case insensitive
	//m - multi line model
	//s - dot matches all

	//when we are making flags we need to make sure that the flags should start with ?
	re = regexp.MustCompile(`(?i)go`)

	//Test string
	text := "Golang is going great"
	//Match
	fmt.Println("Match:", re.MatchString(text))

}
