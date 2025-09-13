package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message.")

	//Set prefix
	log.SetPrefix("Info:")
	log.Println("This is an info message")

	//adding log flags
	//we can chain multiple flags to include in our logs
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with only date information")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an Error Message")

	//let's see how we can add our log messages into standard output instead of terminal
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//0666 is read and write permission
	//0777 in addition to read and write we'll have execute
	if err != nil {
		//fatalf will print o/p onto your terminal and it is going to exit the program
		log.Fatalf("Failed to open log file: %v", err)
	}

	//make sure to close the file as soon as you open any file
	defer file.Close()

	//use info logger to add our log statement to our file
	infoLogger1 := log.New(file, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	warnLogger1 := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger := log.New(file, "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message")
	infoLogger1.Println("This is a info message")
	warnLogger1.Println("This is a warn logger")
	errorLogger1.Println("This is a error message")

	//If we want to use some advanced logging features we can make use of some third party packages
	//one such tool is logrus,zap

}

// debug,info,warn,level, we can use custom logging functions to handle differnet levels of logging
// for that we need to declare new variable that are going to store these different logging levels
var (
	infoLogger  = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
)
