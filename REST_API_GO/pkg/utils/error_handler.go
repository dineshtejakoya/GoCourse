package utils

import (
	"fmt"
	"log"
	"os"
)

func ErrorHandler(err error, message string) error {
	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger.Println(message, err)
	return fmt.Errorf(message)
}

//Reason we are not sending error as return value from this function is we want to keep actual error as private
//private means it can be available to us,we can read the error on to our logger,but the client should not get the actual error
//client needs to get generic error, not the technical error
