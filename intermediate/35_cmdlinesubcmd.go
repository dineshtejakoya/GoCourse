package intermediate

import (
	"flag"
	"fmt"
	"os"
)

// Subcommands are secondary commands that extend the functionality of a main command
// they are specified after a main command and are typically used to perform specific actions or operations
// subcommands promote modularity by separating different functionalities into distinct units
// users can invoke specific functionality directly without needing to remember all options all flags
// in order to get the help file for sub command go run
// >> .\35_cmdlinesubcmd.go firstSub --help,
//  note: go run .\35_cmdlinesubcmd.go --help doesn't work in this case as we are using subcommands
/*	below will provide it in help file for go run .\35_cmdlinesubcmd.go --help
	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(stringFlag)
*/
func main() {

	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(stringFlag)

	// var stringFlag string
	// flag.StringVar(&stringFlag,"user","John","Name of user")

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Command Processing Status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")
	//2flags for first command and 1 flag for second cmd
	//we can throw error if we do not encounter any subcommand when this program is executed
	//done by checking length of the arguments

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
		//exiting with status code 1
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		//as it is referencing we are using pointers
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}

}

//In order to run for first sub command
//go run .\35_cmdlinesubcmd.go firstSub -processing=true -bytes=256
//In order to run second sub command
//go run .\35_cmdlinesubcmd.go secondSub -language Dart
//if it is only one flag in the subcommands = is not mandatory
//if we haven't used = when more than one subcommands it'll only update the starting 1, rest are lost somewhere in the path

//If there is a common flag we want to use for multiple sub commands, then we cannot use the standard format like below
//because it associates one flag with one sub command
//	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
//subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)
