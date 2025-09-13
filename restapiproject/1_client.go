package restapiproject

import (
	"fmt"
	"io"
	"net/http"
)

//In Go the net http package provides the tools to make http requests.
//This package allows you to create http clients that can communicate with web servers using various
// http methods like get,post,put , delete etc..,

// The core component for making http requests in go is the http.Client Struct
// this client can be used to send requests and receive responses from the server
func main() {

	// Create a new http client
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	//resp, err := client.Get("https://swapi.dev/api/people/1")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	//we are going to close the response message that we received from the server
	defer resp.Body.Close()

	// Read and print the response body
	//io.ReadAll will stream on demand
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// fmt.Println(body)
	fmt.Println(string(body))
}
