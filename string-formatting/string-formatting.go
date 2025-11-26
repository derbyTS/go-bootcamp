package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("%05d\n", num)
	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	fmt.Println("Hello \n World")
	fmt.Println("Hello \\n World")
	fmt.Println(`Hello \n World`)
}
