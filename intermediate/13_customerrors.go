package intermediate

import (
	"errors"
	"fmt"
)

//give more info about what went wrong

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type customError struct {
	code    int
	message string
	er      error
}

// Error returns the error message.
// Implementing the Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s %v\n", e.code, e.message, e.er)
}

//Functions that returns a custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong",
// 	}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}
