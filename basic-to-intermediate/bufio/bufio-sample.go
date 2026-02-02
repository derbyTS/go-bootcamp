package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello Bufio\n Test"))

	// Reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read %d bytes %s\n", n, data[:n])
	// fmt.Println(string(data))

	fmt.Println("---------------- ReadString ----------------")

	// There is no left to read cause it already consume by Read. Read doesn't care about the line
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println(line)
	reader1 := bufio.NewReader(strings.NewReader("Hello Bufio\n Test"))
	line1, _ := reader1.ReadString('\n')
	line2, _ := reader1.ReadString('\n')
	fmt.Println("Line 1: ", line1)
	fmt.Println("Line 2: ", line2)

	fmt.Println("---------------- Bufio Writer ----------------")
	writer := bufio.NewWriter(os.Stdout)

	data1 := []byte("Hello writer\n")
	n1, err := writer.Write(data1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Write %d bytes\n", n1)
	// Flush the buffer to ensure all data is written in os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	str := "Hello WriteString"
	n1, err = writer.WriteString(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Write %d string\n", n1)
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
	var b strings.Builder
	writer1 := bufio.NewWriter(&b)
	writer1.WriteString("Hello ")
	writer1.WriteString("World")
	writer1.Flush() // Flush pushes the data to string builder
	result := b.String()
	fmt.Println("This is the value of stringBuilder: ", result)
}
