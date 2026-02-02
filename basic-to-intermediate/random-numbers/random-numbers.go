package main

import (
	"fmt"
	randV1 "math/rand"
	randV2 "math/rand/v2"
	"time"
)

func main() {
	fmt.Println(randV1.Intn(101))
	fmt.Println(randV1.Intn(5) + 1) // rand number from 1 to 5
	fmt.Println(randV2.Int())
	// fix seed
	fmt.Println("Fix seed -------------------")
	fixSeed := randV1.New(randV1.NewSource(10))
	fmt.Println(fixSeed.Intn(100))
	// time as a  seed
	fmt.Println("Time as a seed -------------------")
	timeSeed := randV1.New(randV1.NewSource(time.Now().Unix()))
	fmt.Println(timeSeed.Intn(100))
	fmt.Println("Random Float -------------------")
	fmt.Println(randV1.Float64())
	fmt.Println("Simulating Dice -------------------")

	for {
		fmt.Println("Welcome to Dice")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid Choice")
			continue
		}
		if choice == 2 {
			fmt.Println("Bye")
			break
		}

		die1 := randV1.Intn(6) + 1
		die2 := randV1.Intn(6) + 1
		if choice == 1 {
			fmt.Printf("First Die: %d\n", die1)
			fmt.Printf("Second Die: %d\n", die2)
			fmt.Printf("Total: %d\n", die1+die2)
		}
	}
}
