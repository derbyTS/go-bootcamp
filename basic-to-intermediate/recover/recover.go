package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println("Defer before anything")

	q, err := quotient(5.0, 0)
	if err != nil {
		fmt.Printf("This is an error: %v", err)
	}

	fmt.Println(q)
}

// func quotient(nums ...float64) (quo float64, err error) { float64 don't panic like 5.0 / 0.0  →  +Inf because of IEEE-754
func quotient(nums ...int) (quo int, err error) {
	defer func() {
		if r := recover(); r != nil {
			// err = fmt.Errorf("Error in division: %v", r)
			err = fmt.Errorf("%v", r)
		}
	}()

	quo = nums[0] / nums[1]

	return quo, err
}
