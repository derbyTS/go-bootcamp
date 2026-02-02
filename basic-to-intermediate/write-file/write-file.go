package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := []byte("Hello File1\n\n\n\n")

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data written")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello String")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("String data written")
}
