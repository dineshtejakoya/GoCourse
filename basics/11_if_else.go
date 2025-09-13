package basics

import "fmt"

func main() {
	age := 25

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	temperature := 25
	if temperature >= 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's cool outside")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 are even")
	}

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Both 10 and 5 are even")
	}

}
