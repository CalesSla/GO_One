package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now())

	// specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println(specificTime)

	// parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
	// parsedTime1, _ := time.Parse("06-01-02", "20-05-01")
	// fmt.Println(parsedTime)
	// fmt.Println(parsedTime1)

	t := time.Now()
	// fmt.Println("Formatted time: ", t.Format("06-01-02 04-15"))

	// oneDayLater := t.Add(time.Hour * 24)
	// fmt.Println("One day later: ", oneDayLater)
	// fmt.Println(oneDayLater.Weekday())

	// fmt.Println("Rounded time: ", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)

	// tLocal := t.In(loc)

	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Original time (UTC): ", t)
	// fmt.Println("Original time (Local): ", tLocal)
	// fmt.Println("Rounded time (UTC): ", roundedTime)
	// fmt.Println("Rounded time (Local): ", roundedTimeLocal)

	fmt.Println("Truncated time: ", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("America/New_York")

	tInNY := time.Now().In(loc)
	fmt.Println("New Yourk time: ", tInNY)

	t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)

	duration := t2.Sub(t1)
	fmt.Println("Duration: ", duration)

	fmt.Println("t2 is after t1: ", t2.After(t1))

}
