package main

import "fmt"

func main() {
	greeting := make(chan string)
	greetingHello := "Hello World"

	go func() {
		greeting <- greetingHello
		greeting <- "Hello Again"
		for _, v := range "abcde" {
			greeting <- string(v)
		}
	}()

	go func() {
		receiver := <-greeting
		fmt.Println(receiver)
	}()

	receiver := <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	fmt.Println("End")
}
