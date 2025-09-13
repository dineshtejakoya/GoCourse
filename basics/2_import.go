package basics

import (
	"fmt"
	foo "net/http"
	//renaming net/http import statement to foo
)

func main() {
	fmt.Println("Hello, Go Standard Library")

	//replacing http.Get to foo.Get
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response Status:", resp.Status)
}

/*
Note while running above code below error came up:
Hello, Go Standard Library
Error:  Get "https://jsonplaceholder.typicode.com/posts/1": dial tcp: lookup jsonplaceholder.typicode.com: no such host
this would usually happen due to internet or dns resolution issue
Resolution: open command prompt as administrator and run : ipconfig/flushdns
it should resolve
*/
