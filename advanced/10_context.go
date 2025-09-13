package advanced

import (
	"context"
	"fmt"
)

//contexts can carry key value pairs
//These are instances of struct or Objects
/*
Context is a type from the context package
contexts are used to carry deadlines, cancellation signals and request scoped values
Apart from that we are carrying values in key value pairs
Contexts are closely associated with APIs
contexts are also used in goroutines to carry values

*/

/*
Best Practices of context:
Avoid storing contexts in struct fields or global variables
use context values for request scope data, not for passing optional parameters
ensure context keys are unexported to prevent collisions
When it comes to long running operations always check for context cancellation, ensure to cleanup resources and handle errors properly when a context is cancelled
creating context inside loops lead to resource leaks
*/

// Difference between context.TODO() and context.Background()
func main() {
	//context.TODO() is used when you are unsure about which context to use or if you plan to use a proper context later
	//it acts as a placeholder and it doesn't carry any deadlines and cancellations
	todoContext := context.TODO()
	//Apart from context.TODO() we can also use context.Background() to create context
	contextBkg := context.Background()
	ctx := context.WithValue(todoContext, "name", "John")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx1 := context.WithValue(contextBkg, "city", "NewYork")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value("city"))

}
