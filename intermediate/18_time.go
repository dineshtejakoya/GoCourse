package intermediate

import (
	"fmt"
	"time"
)

func main() {

	//Current Local Time
	fmt.Println(time.Now())

	//Specific time
	specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	//Parse Time
	//layout is for very specific date
	//type is time.Time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")
	//it won't give results as expected
	//because go uses a very specific reference date
	//i.e., Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-01-02 14-04", "20-05-01 18-03")
	parsedTime3, _ := time.Parse("06-01-02 15-04", "20-05-01 18-03")
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	//Formatting Time
	t := time.Now()
	//it referencing same value, in below case minutes are added before hours
	fmt.Println("Formatted Time:", t.Format("06-01-02 04-15"))

	//add one day
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	//it is confusing? pls check later
	fmt.Println("Rounded Time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)

	//Convert this to the specific time zone
	tLocal := t.In(loc)

	//Perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original Time (UTC)", t)
	fmt.Println("Original Time (Local):", tLocal)
	fmt.Println("Rounded Time (UTC):", roundedTime)
	fmt.Println("Rounded Time (Local):", roundedTimeLocal)

	//Truncation
	//truncate will always round the clock down when giving input
	fmt.Println("Truncated Time:", t.Truncate(time.Hour))

	loc1, _ := time.LoadLocation("America/New_York")
	//convert  time to location
	tInNY := time.Now().In(loc1)
	fmt.Println("New York Time:", tInNY)
}
