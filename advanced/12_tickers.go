package advanced

import (
	"fmt"
	"time"
)

// Tickers in Go is a mechanism for producing ticks at regular intervals
// They are useful for performing periodic tasks or operations on a consistent schedule
// often used in scenarios where tasks need to be repeated at fixed intervals such as polling, periodic logging or regular updates
func main() {
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()
	//tick at each second
	// for tick := range ticker.C {
	// 	fmt.Println("Tick at:", tick)
	// }

	i := 1
	for range 10 {
		i *= 2
		fmt.Println(i)
	}

	for tick := range ticker.C {
		i *= 2
		fmt.Println(tick)
	}
}

//To stop a ticker from releasing resources and prevent it from producing further ticks,
//because ticker does not expires, timer have an expiry but tickers don't
//that's why we always have to stop a ticker
//use defer to stop the ticker before the end of the functions returned
