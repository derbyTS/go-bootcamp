package main

import "fmt"

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
