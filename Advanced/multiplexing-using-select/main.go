package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		time.Sleep(time.Second)
		ch <- 1
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel is close")
				return
			} else {
				fmt.Println(msg)
			}
		default:
			fmt.Println("No channels ready....")
		}
	}
}

// Channel timeout
// func main() {
// 	ch := make(chan int)
//
// 	go func() {
// 		defer close(ch)
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 	}()
// 	time.Sleep(2 * time.Second)
// 	select {
// 	case msg, ok := <-ch:
// 		if !ok {
// 			fmt.Println("Channel Close")
// 		}
// 		fmt.Println(msg)
// 	case <-time.After(time.Second):
// 		fmt.Println("Timeout")
// 	}
// }

// Multiple channel
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
//
// 	go func() {
// 		// time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()
//
// 	go func() {
// 		// time.Sleep(time.Second)
// 		ch2 <- 2
// 	}()
//
// 	time.Sleep(2 * time.Second)
// 	for range 2 {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("Received from ch1: ....", msg1)
//
// 		case msg2 := <-ch2:
// 			fmt.Println("Received from ch2: ....", msg2)
// 		default:
// 			fmt.Println("No channels ready....")
// 		}
// 	}
// }
