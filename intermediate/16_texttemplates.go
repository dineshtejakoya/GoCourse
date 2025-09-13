package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	// //using text/template package
	// tmpl := template.New("example")

	// //now once we have created the template
	// //we have to parse/process the template using Parse() function
	// tmpl, err := tmpl.Parse("Welcome, {{.name}}! How are you doing?\n")
	// if err != nil {
	// 	panic(err)
	// }

	// //Define data for the welcome message template
	// data := map[string]interface{}{
	// 	"name": "John",
	// }

	// //Now, lets render the welcome message out to the console
	// err = tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	// //Another way we  initiate our template i.e., template.Must
	// //When we use template.must we'll not have to handle the template ourselves and template will automatically panic if we have an error while parsing our template
	// tmpl1 := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))
	// //.must is used to handle from parse

	// //Define data for the welcome message template
	// data1 := map[string]interface{}{
	// 	"name": "Dinesh",
	// }

	// //Now, lets render the welcome message out to the console
	// err = tmpl1.Execute(os.Stdout, data1)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	//\n is a delimiter we are using now, which means it is going to accept the input from the console until it reads new line character
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Define named templates for different types of
	//the name of the template and the content of the template are going to be string
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}} we are glad you joined",
		"notification": "{{.nm}}, you have a new notification: {{.ntf}}",
		"error":        "Oops! An error occured: {{.em}}",
	}

	//Parse and store templates
	//It is a slice that is going to store all the templates
	//string is name of the map
	//key is template, denoted by *template.Template
	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl2 := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl2))
	}

	for {
		// Show Menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4.Exit")
		fmt.Println("Choose an Option:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl2 *template.Template

		switch choice {
		case "1":
			tmpl2 = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notification message")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl2 = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "ntf": notification}

		case "3":
			fmt.Println("Enter your error message:")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl2 = parsedTemplates["error"]
			data = map[string]interface{}{"nm": name, "em": errorMessage}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Choice. Please select a valid option.")
			continue
		}

		//render and print the template to the console
		err := tmpl2.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	}
}
