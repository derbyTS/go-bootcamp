package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	// This is added by gemini: The idea is to use --help without the sub-command
	// Define what the "Global" help looks like
	flag.Usage = func() {
		fmt.Println("Usage: mytool [command] [flags]")
		fmt.Println("\nCommands:")
		fmt.Println("  firstSub   Run the first process")
		fmt.Println("  secondSub  Run the second process")
	}
	// You MUST call Parse to catch the top-level --help
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("This program needs additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing: ", *firstFlag)
		fmt.Println("bytes: ", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language: ", *flagsc2)
	default:
		fmt.Println("No sub-command entered")
		os.Exit(1)
	}
}
