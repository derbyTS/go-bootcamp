package main

import (
	"fmt"
	"time"
	localLocation "time"
)

func main() {
	fmt.Println(localLocation.Now())
	specificTime := localLocation.Date(
		1993,
		localLocation.August,
		27,
		8,
		30,
		30,
		30,
		localLocation.UTC,
	)
	fmt.Println(specificTime)

	// Reference Time in Go - Mon Jan 2 15:04:05 MST 2006
	// Parse Time
	parsedTime, _ := localLocation.Parse("2006-01-02", "2020-05-01")
	parsedTime1, _ := localLocation.Parse("06-01-02", "20-05-01")
	parsedTime2, _ := localLocation.Parse("06-1-2", "20-5-1")
	parsedTime3, _ := localLocation.Parse("06-1-2 15-04", "20-5-1 18-03")

	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)

	// Fomatting time
	t := localLocation.Now()
	fmt.Println(t)
	fmt.Println("Formatted Time ", t.Format("06-01-02 15:04:05"))
	fmt.Println("Formatted Time ", t.Format("Monday 06-01-02 15:04:05"))
	oneDayLater := t.Add(localLocation.Hour * 24)
	fmt.Println("One day later ", oneDayLater.Format("Monday 06-01-02 15:04:05"))
	fmt.Println("One day later day", oneDayLater.Weekday())
	fmt.Println("Round Time Local Time", t.Round(localLocation.Hour))
	fmt.Println(
		"Truncate Time Local Time",
		t.Round(localLocation.Hour),
	) // Round can go up and down but truncate always go down

	locTime, _ := localLocation.LoadLocation("Asia/Manila")
	bday := localLocation.Date(2026, localLocation.September, 23, 00, 01, 02, 02, locTime)
	fmt.Println(bday.Weekday())

	// Convert to the specific timezone
	ph := localLocation.Date(
		2026,
		localLocation.September,
		23,
		00,
		01,
		02,
		02,
		localLocation.UTC,
	)
	madrid, _ := localLocation.LoadLocation("Europe/Madrid")
	newYork, _ := localLocation.LoadLocation("America/New_York")

	roundTimeUtc := bday.Round(localLocation.Hour)
	roundTimeLocal := roundTimeUtc.In(locTime)
	roundTimeMadrid := roundTimeUtc.In(madrid)
	roundTimeNewYork := roundTimeUtc.In(newYork)

	fmt.Println("Local time: ", bday)
	fmt.Println("UTC time: ", ph)
	fmt.Println("Round time Manila: ", roundTimeLocal)
	fmt.Println("Round time Madrid", roundTimeMadrid)
	fmt.Println("Round time New York", roundTimeNewYork)

	// Convert the time in local time
	newYorkTime := time.Now().In(newYork)
	fmt.Println("Current New York time now base on my OS clock is ", newYorkTime)

	// Subtract Time
	fmt.Println("-------------------subtract time---------------------")
	tNo1 := localLocation.Now()
	lp()
	tNo2 := localLocation.Now()
	fmt.Println(tNo2.Sub(tNo1))
}

func lp() {
	for i := 0; i < 100000000; i++ {
	}
}
