package restapiproject

import (
	"fmt"
	"log"
	"net/http"
)

//Note: localhost or 127.0.0.1 both the same thing

//To create http server in go we need to define http handlers, which are functions
//They simply put functions that handle the busienss logic for an endpoint
//handler will take care what should happen when client sent a request to the endpoint

// We also configure routes
func main() {

	//we need to configure
	//http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request))
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server")
	})

	//127.0.0.0 is an address for localhost, not an external ip address
	//3000 is the port, request and response is sent through port 3000
	const serverAddr string = "127.0.0.1:3000"

	fmt.Println("Server Listening on port 3000")
	//listen and serve the request
	//we need handle erroralso here,
	// if there is a technical problem, we already have something that is running on the port then we receive an error message that this port is not avaialble
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}

}

//instead of using address "127.0.0.1:3000" we can delete "127.0.0.1" and use the port itself
//then it will understand we are using localhost
//note: the ip address of localhost is "127.0.0.1"
