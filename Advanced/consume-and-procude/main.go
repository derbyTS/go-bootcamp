package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Unbufferecd")
	startUnbuffered := time.Now()
	ch := make(chan int) // Unbufferecd channel
	go produce(ch)
	consume(ch)
	fmt.Printf("Unbuffered Total Time: %v\n\n", time.Since(startUnbuffered))

	startBuffered := time.Now()
	fmt.Println("Buffered")
	ch1 := make(chan int, 3) // Buffered channel
	go produce(ch1)
	consume(ch1)
	fmt.Printf("Buffered Total Time: %v\n\n", time.Since(startBuffered))
}

func produce(ch chan<- int) {
	defer close(ch)
	for i := range 3 {
		ch <- i
		fmt.Println("Producing message: ", i)
	}
}

func consume(ch <-chan int) {
	for val := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println("Consuming message: ", val)
	}
}
