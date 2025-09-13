package intermediate

import (
	"flag"
	"fmt"
	"os"
)

// In Go Command line args are accessible through os.args slice
// where os.args[0] is the name of the command or the program name itself
// flag package modify the behavior of the program
func main() {

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("Argument:", i, ":", arg)
	}

	//Define Flags
	var name string
	var age int
	var male bool

	//default value is john, and description is name of the user
	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
	//in order to pass through cmd line parameters go run .\34_cmdlineargs.go -name Dinesh Teja -age 25
	//in the above example even though we pass Dinesh Teja , the o/p we receive is Dinesh
	//coz as soon as we encountered space it's a different argument
	//in order to accept complete argument use "" in between them, now the arguments count is also decreased
	//if we are not using flags(-name,-age,-male), it will show the default values
}

//When we have command line flags with their usage description available for the main command, that is, the main program, we can use
//--help to get them all printed on the terminal, to get to know more about the program/command and different options available to us with the command.
//go run 34_cmdlinears.go --help will provide the details of description of flags
