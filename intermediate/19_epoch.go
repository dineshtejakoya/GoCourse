package intermediate

import (
	"fmt"
	"time"
)

// epoch refers to specific point in time that servers as reference for timestamps and calculations
//
//00:00:00 UTC on jan 1, 1970
func main() {
	//unix time does not account for leap seconds which are adjustments made to keep time synchronized with earths rotation, which can lead to sight inaccuracies over longer period
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime) //unreadable big number 1754278006
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time:", t.Format("2006-01-02"))

}
