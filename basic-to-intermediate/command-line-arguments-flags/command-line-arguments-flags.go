package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command: ", os.Args[0])
	for i, v := range os.Args {
		fmt.Printf("Args %d: %s\n", i, v)
	}

	// Define flags
	var name string
	var age int
	var isSenior bool

	flag.StringVar(&name, "name", "Derby", "Name of the user")
	flag.IntVar(&age, "age", 0, "Age of the user")
	flag.BoolVar(&isSenior, "isSenior", false, "Is the user senior")

	flag.Parse()

	fmt.Println("Name: ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Senior: ", isSenior)
}
