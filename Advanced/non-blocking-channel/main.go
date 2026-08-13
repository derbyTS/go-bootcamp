package main

import (
	"fmt"
	"time"
)

func main() {
	// ==== Non-blocking received operation
	// ch := make(chan int)

	// select {
	// case msg := <-ch:
	// 	fmt.Println(msg)
	// default:
	// 	fmt.Println("No message")
	// }

	// // ==== Non-blocking received operation
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent message")
	// default:
	// 	fmt.Println("Channel is not ready to receive")
	// }

	// === Non-blocking operation in real time system
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received: ", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	fmt.Println("Before Sending")
	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
}
