package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 seconds after")
	}()

	go func() {
		ch <- 2
		time.Sleep(3 * time.Second)
		fmt.Println("3 seconds after")
	}()

	receiver := <-ch
	fmt.Println(receiver)
	fmt.Println(<-ch)
	fmt.Println("End")
}
