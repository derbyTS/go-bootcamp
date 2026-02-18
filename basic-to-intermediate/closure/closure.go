package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence1 := adder()
	fmt.Println(sequence1())

	adderMain := func() func(int) int {
		return func(i int) int {
			i++
			return i
		}
	}()

	fmt.Println(adderMain(1))
	fmt.Println(adderMain(1))
	fmt.Println(adderMain(1))
	fmt.Println(adderMain(1))
	fmt.Println(adderMain(1))

	adderMainWithValue := func() func(int) int {
		x := 0
		return func(i int) int {
			x += i
			return x
		}
	}()

	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	fmt.Println(adderMainWithValue(1))
	prepSum := PrepareSum(5, 2)
	fmt.Println(prepSum())
	fmt.Println(prepSum())
}

func adder() func() int {
	i := 0
	fmt.Printf("Prev value of i: %d\n", i)
	return func() int {
		i++
		fmt.Printf("After adding to value of i: %d\n", i)
		return i
	}
}

// anonymous function
func PrepareSum(a int, b int) func() string {
	// We define a "closure" here.
	// This anonymous function 'remembers' a and b.
	return func() string {
		sum := a + b
		return fmt.Sprintf("The total sum is: %d", sum)
	}
}
