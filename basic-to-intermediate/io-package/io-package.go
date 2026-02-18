package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func closerResouce(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer")
	fmt.Println(buf.String())
}

func multiReadExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(mr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer closerResouce(file)

	// _, err = file.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }

	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader"))

	fmt.Println("=== Write from Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader Example ===")
	multiReadExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	fmt.Println("=== Creating file Example ===")
	writeToFile("io.txt", "Hello Writer")

	fmt.Println("=== Closing Resource Example ===")
	myResource := &MyResource{name: "Test Resource"}
	closerResouce(myResource)
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing: ", m)
	return nil
}
