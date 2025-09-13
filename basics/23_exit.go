package basics

import (
	"fmt"
	"os"
)

//OS.Exit() is a function that terminates the program immediately with the given status code
//It's useful for situations where you need to halt the execution of the program completely,
// without performing any defer functions or performing any cleanup operations

//status codes : 0 is success, non zero is failure

//os.Exit() is used in unrecoverable error scenarios or when the program must stop immediately

func main() {

	defer fmt.Println("Deferred Statement")
	//when os.exit() is called it stops the program execution immediately,
	//  and any deferred functions registered using defer will not be executed
	fmt.Println("Starting the main function")

	//Exit with status code of 1
	os.Exit(1)

	//This will never be executed
	fmt.Println("End of main function")

}
