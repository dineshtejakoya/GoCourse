package intermediate

import (
	"fmt"
	"net/url"
)

// parse simply means we are processing some data and we are extracting some info from that data
func main() {

	//[scheme://][userinfo@]host[:port][/path][?query][#fragment]
	//scheme called protocol which can be http,https/ftp
	//userinfo contains username and pwd which is optional
	//host ->domain name or ip address
	//port optional
	//path is resources on to the servers
	//fragment identifier and used to specify location for resources

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error Parsing URL:", err)
		return
	}

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Port: ", parsedURL.Port())
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)
	fmt.Println("Fragment: ", parsedURL.Fragment)

	fmt.Println("==========Query Parameters===========")
	rawURL1 := "https://example.com/path?name=John&age=30"

	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	//Building URL
	fmt.Println("=======Building URL==========")
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL:", baseURL.String())

	fmt.Println("========Creating Query & Adding Query to URL========")
	values := url.Values{}
	//Add key-value pairs to the values object
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	//Encode
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	//Now all we need to do add this query to any url
	//Build a URL
	baseURL1 := "https://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery
	fmt.Println(fullURL)

}
