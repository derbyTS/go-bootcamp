package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		fmt.Println("Closing the file")
		file.Close()
	}()

	fmt.Println("File is open")

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The string value from file is ", string(data))
	fmt.Printf("The byte values from file is %v\n", data)

	file1, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		fmt.Println("Closing file1")
		file1.Close()
	}()

	scanner := bufio.NewScanner(file1)
	fmt.Printf("%+v\n", scanner)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}
