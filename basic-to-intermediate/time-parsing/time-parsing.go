package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02T15:04:05Z07:00"
	str := "2025-12-08T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Thre is an error: ", err)
		return
	}
	fmt.Println("The date is : ", t)

	str1 := "Dec 08, 2025 06:21 PM"
	layout1 := "Jan 02, 2006 03:04 PM"

	t1, err := time.Parse(layout1, str1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The date is : ", t1)
}
