package intermediate

//because we are using logrus , we need to import our package
//for that we need to initialize our module
//>>go mod init gocourse
//go.mod file is going to contain all the dependencies that are related to our current package
import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	//set log level
	log.SetLevel(logrus.InfoLevel)

	//Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	//Logging Examples
	log.Info("This is an info message")
	log.Warn("This is a warn message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("Usser logged in.")
}
