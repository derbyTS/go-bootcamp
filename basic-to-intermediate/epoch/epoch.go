package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current unix time: ", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Date: ", t.Format("2006-01-02"))
}
