package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in initializing zap logger")
	}
	//zap logger may contain some buffer by the end of the function
	//it's better to flush the buffer that the logger contains

	defer logger.Sync()

	logger.Info("This is an info message..")

	logger.Info("User logged in", zap.String("username", "john doe"), zap.String("method", "GET"))

}
